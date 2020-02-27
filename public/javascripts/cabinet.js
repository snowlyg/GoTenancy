'use strict';
import '../stylesheets/home_banner.scss'
import '../stylesheets/home_products.scss'
import '../stylesheets/qor.scss'
import '../stylesheets/qor_auth.scss'

$(function () {
    /*   $("#add-cashe").submit(function(event) {
        event.preventDefault();
        $.ajax({
          type: "POST",
          url: "/cabinet/add_user_credit",
          error: function(xhr) {
            alert(xhr.status + ": " + xhr.statusText);
          },
          success: function(response) {
            alert(response.status + ": " + response.message + " (" + response.itemID + ")");
            location.reload();
          },
          data: $(event.target).serialize()
        });
      }); */

    /*   $("#user-profile").submit(function(event) {
        event.preventDefault();
        $.ajax({
          type: "POST",
          url: "/cabinet/profile",
          error: function(xhr) {
            alert(xhr.status + ": " + xhr.statusText);
          },
          success: function(response) {
            console.log(response.status + ": " + response.message + " (" + response.itemID + ")");
            // location.reload();
          },
          data: $(event.target).serialize()
        });
      });
      $("#billing-address").submit(function(event) {
        event.preventDefault();
        $.ajax({
          type: "POST",
          url: "/cabinet/profile/billing_address",
          error: function(xhr) {
            alert(xhr.status + ": " + xhr.statusText);
          },
          success: function(response) {
            alert(response.status + ": " + response.message + " (" + response.itemID + ")");
            // location.reload();
          },
          data: $(event.target).serialize()
        });
      });
      $("#shipping-address").submit(function(event) {
        event.preventDefault();
        $.ajax({
          type: "POST",
          url: "/cabinet/profile/shipping_address",
          error: function(xhr) {
            alert(xhr.status + ": " + xhr.statusText);
          },
          success: function(response) {
            alert(response.status + ": " + response.message + " (" + response.itemID + ")");
            // location.reload();
          },
          data: $(event.target).serialize()
        });
      }); */
})
