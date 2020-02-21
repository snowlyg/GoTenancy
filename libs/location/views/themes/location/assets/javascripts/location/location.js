'use strict';

function qorInitMapAssets(selector) {
    $(selector).each(function () {
        let $selector = $(this),
            mapData = $selector.data(),
            mapSource = mapData.locationMapsource,
            apiKey = mapData[`${mapSource}Apikey`],
            id = `qor_map_${mapSource}`,
            src;

        if (mapSource === 'baidu') {
            if ($(`#${id}`).length) {
                new QorLocation($(this)).initBaiduMap();
                return;
            }
            src = `//api.map.baidu.com/api?v=2.0&ak=${apiKey}&callback=qorInitBaiduMap`;
        } else if (mapSource === 'google') {
            if ($(`#${id}`).length) {
                new QorLocation($(this)).initGoogleMap();
                return;
            }
            src = `//maps.googleapis.com/maps/api/js?key=${apiKey}&callback=qorInitGoogleMap`;
        } else {
            console.log('QOR only support baidu or google Map, please make sure your Map config is correct!');
        }

        qorLoadMapScript(src, id);
    });
}

function qorInitBaiduMap() {
    $('[data-toggle="qor.location"]').each(function () {
        new QorLocation($(this)).initBaiduMap();
    });
}

function qorInitGoogleMap() {
    $('[data-toggle="qor.location"]').each(function () {
        new QorLocation($(this)).initGoogleMap();
    });
}

function qorLoadMapScript(src, id) {
    let script = document.createElement('script');
    script.src = src;
    script.id = id;
    document.body.appendChild(script);
}

function QorLocation(element, options) {
    this.$element = $(element);
    this.options = $.extend({}, QorLocation.DEFAULTS, $.isPlainObject(options) && options);
    this.roles = {};
    this.init();
}

QorLocation.prototype = {
    constructor: QorLocation,

    init: function () {
        let $element = this.$element,
            $roles = $element.find('[data-location-role]'),
            roles = this.roles,
            roleName;

        this.MAP_ID = $element.find('.qor__location-map')[0];
        this.isGoogle = $element.data('location-mapsource') === 'google';
        this.$location = $element.find('.qor__location-address');

        // country city region address zip latitude longitude geocode reverseGeocode map currentAddress
        $roles.each(function () {
            roleName = $(this).data('location-role');
            roles[`$${roleName}`] = $(this);
        });
        this.bind();
    },

    bind: function () {
        this.roles.$reverseGeocode.on('click', this.getAddress.bind(this));
        this.roles.$geocode.on('click', this.getPosition.bind(this));
    },

    initGoogleMap: function () {
        let options = {
                enableHighAccuracy: true,
                timeout: 5000,
                maximumAge: 0
            },
            roles = this.roles,
            latitude = Number(roles.$latitude.val()),
            longitude = Number(roles.$longitude.val()),
            _this = this;

        if (latitude === 0 && longitude === 0) {
            navigator.geolocation.getCurrentPosition(
                function (pos) {
                    _this.setupGoogleMap({lat: pos.coords.latitude, lng: pos.coords.longitude});
                },
                function (err) {
                    console.warn(`ERROR(${err.code}): ${err.message}`);
                    return;
                },
                options
            );
        } else {
            this.setupGoogleMap({lat: latitude, lng: longitude});
        }
    },

    setupGoogleMap: function (point) {
        this.map = new google.maps.Map(this.MAP_ID, {
            center: point,
            zoom: 16
        });
        this.marker = new google.maps.Marker({
            position: point,
            draggable: true,
            map: this.map
        });

        this.mapGeo = new google.maps.Geocoder();
        google.maps.event.bind(this.marker, 'dragend', this, this.moveMarker);
        this.moveMarker();
    },

    initBaiduMap: function () {
        let roles = this.roles,
            latitude = Number(roles.$latitude.val()),
            longitude = Number(roles.$longitude.val()),
            point = new BMap.Point(longitude, latitude),
            _this = this;

        if (latitude === 0 && longitude === 0) {
            let geolocation = new BMap.Geolocation();
            geolocation.getCurrentPosition(
                function (r) {
                    if (this.getStatus() == BMAP_STATUS_SUCCESS) {
                        _this.setupBaiduMap(r.point);
                    } else {
                        console.log('failed' + this.getStatus());
                    }
                },
                {enableHighAccuracy: true}
            );
        } else {
            this.setupBaiduMap(point);
        }
    },

    setupBaiduMap: function (point) {
        let map = new BMap.Map(this.MAP_ID),
            marker = new BMap.Marker(point);

        map.centerAndZoom(point, 18);
        map.enableScrollWheelZoom();

        map.addOverlay(marker);
        map.panTo(point);
        marker.enableDragging();
        marker.addEventListener('dragend', this.moveMarker.bind(this));

        this.marker = marker;
        this.map = map;
        this.mapGeo = new BMap.Geocoder();
        this.moveMarker();
    },

    moveMarker: function () {
        let $location = this.$location,
            mapGeo = this.mapGeo,
            pos = this.marker.getPosition();

        if (this.isGoogle) {
            mapGeo.geocode({latLng: pos}, function (results, status) {
                if (status == google.maps.GeocoderStatus.OK) {
                    $location.html(results[0].formatted_address);
                }
            });
        } else {
            mapGeo.getLocation(pos, function (result) {
                if (result) {
                    $location.html(result.address);
                }
            });
        }
    },

    getPosition: function () {
        let map = this.map,
            roles = this.roles,
            marker = this.marker,
            $location = this.$location,
            city = roles.$city.val(),
            address = roles.$address.val(),
            myGeo = this.mapGeo;

        if (this.isGoogle) {
            myGeo.geocode({address: `${address}, ${city}`}, function (results, status) {
                if (status === google.maps.GeocoderStatus.OK) {
                    let result = results[0],
                        pos = result.geometry.location;

                    map.setCenter(pos);
                    marker.setPosition(pos);
                    roles.$latitude.val(pos.lat);
                    roles.$longitude.val(pos.lng);
                    $location.html(result.formatted_address);
                } else {
                    QOR.qorConfirm('Geocode was not successful for the following reason: ' + status);
                }
            });
        } else {
            if (city == '' || address == '') {
                QOR.qorConfirm('Please ensure that the city and address have been filled in correctly!');
                return;
            }
            myGeo.getPoint(
                address,
                function (point) {
                    if (point) {
                        roles.$latitude.val(point.lat);
                        roles.$longitude.val(point.lng);
                        map.centerAndZoom(point, 18);
                        marker.setPosition(point);
                    }
                },
                city
            );
        }
    },

    getAddress: function () {
        let pos = this.marker.getPosition(),
            myGeo = this.mapGeo,
            roles = this.roles,
            _this = this;

        roles.$latitude.val(pos.lat);
        roles.$longitude.val(pos.lng);

        if (this.isGoogle) {
            myGeo.geocode({location: pos}, function (results, status) {
                if (status == google.maps.GeocoderStatus.OK) {
                    _this.setLocationFromGoogle(results[0]);
                }
            });
        } else {
            myGeo.getLocation(new BMap.Point(pos.lng, pos.lat), function (result) {
                if (result) {
                    roles.$address.val(result.address);
                    roles.$city.val(result.addressComponents.city);
                    roles.$region.val(result.addressComponents.district);
                }
            });
        }
    },

    setLocationFromGoogle: function (data) {
        let roles = this.roles,
            country = '',
            address = '',
            city = '',
            region = '',
            zip = '',
            values = {
                street_number: '',
                route: '',
                postal_code: '',
                administrative_area_level_1: '',
                locality: '',
                country: ''
            },
            divided_by_city,
            left_side,
            addressComponents = data['address_components'];

        for (var compIndex in addressComponents) {
            var types = addressComponents[compIndex].types;
            for (var typeIndex in types) {
                if (types[typeIndex] == 'country') {
                    country = addressComponents[compIndex].long_name;
                }
                values[types[typeIndex]] = addressComponents[compIndex].short_name;
            }
        }

        city = values['locality'] || values['sublocality'] || values['administrative_area_level_3'];
        region = values['country'] == 'FR' ? values['administrative_area_level_1'] : values['administrative_area_level_2'] || values['administrative_area_level_1'];
        zip = values['postal_code'];

        divided_by_city = data.formatted_address.split(city);
        left_side = divided_by_city[0];

        if (zip && zip != '') {
            left_side = left_side.replace(zip, '');
        }
        if (left_side.replace(region, '') == left_side) {
            address = left_side;
        } else {
            address = divided_by_city[1];
        }
        address = address.replace(/[\,\s]+$/i, '');
        if (data.formatted_address.match(/^\d*,/)) address = '';

        roles.$country.val(country);
        roles.$zip.val(zip);
        roles.$region.val(region);
        roles.$city.val(city);
        roles.$address.val(address);
    }
};

$(function () {
    let selector = '[data-toggle="qor.location"]',
        NAMESPACE = 'qor.location',
        EVENT_ENABLE = 'enable.' + NAMESPACE;

    qorInitMapAssets(selector);

    $(document).on(EVENT_ENABLE, function () {
        qorInitMapAssets(selector);
    });
});
