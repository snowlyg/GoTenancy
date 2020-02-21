package seo

// MicroProductTemplate a bundle of microdata templates
var (
	MicroProductTemplate = `
	<div itemscope itemtype="http://schema.org/Product" style="display:none;">
  <span itemprop="brand">{{.BrandName}}</span>
  <span itemprop="name">{{.Name}}</span>
  <img itemprop="image" src="{{.Image}}" />
  <span itemprop="description">{{.Description}}</span>
  <span itemprop="sku">{{.SKU}}</span>
  <span itemprop="aggregateRating" itemscope itemtype="http://schema.org/AggregateRating">
    <span itemprop="ratingValue">{{.RatingValue}}</span> <span itemprop="reviewCount">{{.ReviewCount}}</span>
  </span>

  <span itemprop="offers" itemscope itemtype="http://schema.org/Offer">
    <meta itemprop="priceCurrency" content="USD" />
    <span itemprop="price">{{.Price}}</span>
    <time itemprop="priceValidUntil" datetime="{{.PriceValidUntil}}"></time>
    <span itemprop="seller">{{.SellerName}}</span>
      <link itemprop="itemCondition" href="http://schema.org/UsedCondition"/>
      <link itemprop="availability" href="http://schema.org/InStock"/>
    </span>
  </span>
	</div>
	`

	MicroContactTemplate = `
	<script type="application/ld+json">
	{ "@context" : "http://schema.org",
		"@type" : "Organization",
		"url" : "{{.URL}}",
		"contactPoint" : [
			{ "@type" : "ContactPoint",
				"telephone" : "{{.Telephone}}",
				"contactType" : "{{.ContactType}}"
			} ] }
	</script>
	`

	MicroSearchTemplate = `
	<script type="application/ld+json">
	{
		"@context": "http://schema.org",
		"@type": "WebSite",
		"url": "{{.URL}}",
		"potentialAction": {
			"@type": "SearchAction",
			"target": "{{.Target}}",
			"query-input": "{{.FormattedQueryInput}}"
		}
	}
	</script>
	`
)
