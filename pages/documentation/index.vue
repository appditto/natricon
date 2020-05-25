<template>
  <div class="w-full flex flex-col items-center">
    <documentation-hero />
    <div class="n-container px-6 md:px-12">
      <divider />
    </div>
    <endpoint />
    <div class="n-container px-6 md:px-12">
      <divider />
    </div>
    <required-parameters />
    <div class="n-container px-6 md:px-12">
      <divider />
    </div>
    <optional-parameters />
    <div class="n-container px-6 md:px-12">
      <divider />
    </div>
  </div>
</template>
<style>
pre {
  white-space: pre-wrap !important; /* Since CSS 2.1 */
  white-space: -moz-pre-wrap !important; /* Mozilla, since 1999 */
  white-space: -pre-wrap !important; /* Opera 4-6 */
  white-space: -o-pre-wrap !important; /* Opera 7 */
  word-wrap: break-word !important; /* Internet Explorer 5.5+ */
}
</style>
<script>
import DocumentationHero from "~/components/DocumentationHero.vue";
import Divider from "~/components/Divider.vue";
import Endpoint from "~/components/Endpoint.vue";
import RequiredParameters from "~/components/RequiredParameters.vue";
import OptionalParameters from "~/components/OptionalParameters.vue";
export default {
  components: {
    DocumentationHero,
    Divider,
    Endpoint,
    RequiredParameters,
    OptionalParameters
  },
  data() {
    return {
      pageDescription: "learn how to integrate natricon into nano services.",
      pageTitle: "natricon | api documentation",
      pagePreview:
        "https://natricon.com/images/previews/preview-documentation.png",
      pageThemeColor: "#FFFFFF",
      canonicalURL: "https://natricon.com/documentation"
    };
  },
  head() {
    return {
      htmlAttrs: {
        lang: "en"
      },
      title: this.pageTitle,
      meta: [
        // hid is used as unique identifier. Do not use `vmid` for it as it will not work
        {
          hid: "description",
          name: "description",
          content: this.pageDescription
        },
        // Google / Search Engine Tags
        {
          itemprop: "name",
          content: this.pageTitle
        },
        {
          itemprop: "description",
          content: this.pageDescription
        },
        {
          itemprop: "image",
          content: this.pagePreview
        },
        // Facebook Meta Tags
        {
          property: "og:url",
          content: this.canonicalURL
        },
        {
          property: "og:type",
          content: "website"
        },
        {
          property: "og:title",
          content: this.pageTitle
        },
        {
          property: "og:description",
          content: this.pageDescription
        },
        {
          property: "og:image",
          content: this.pagePreview
        },
        // Twitter Meta Tags
        {
          name: "twitter:card",
          content: "summary_large_image"
        },
        {
          name: "twitter:title",
          content: this.pageTitle
        },
        {
          name: "twitter:description",
          content: this.pageDescription
        },
        {
          name: "twitter:image",
          content: this.pagePreview
        },
        // Theme
        {
          name: "theme-color",
          content: this.pageThemeColor
        },
        // Windows 8 IE 10
        {
          name: "msapplication-TileColor",
          content: this.pageThemeColor
        }
      ],
      link: [
        // Canonical
        {
          rel: "canonical",
          href: this.canonicalURL
        },
        // Generic Icons
        {
          rel: "icon",
          sizes: "180x180",
          href: "/apple-touch-icon.png"
        },
        { rel: "icon", sizes: "32x32", href: "/favicon-32x32.png" },
        { rel: "icon", sizes: "16x16", href: "/favicon-16x16.png" },
        { rel: "manifest", href: "/site.manifest" }
      ]
    };
  }
};
</script>