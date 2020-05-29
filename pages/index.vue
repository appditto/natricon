<template>
  <div class="w-full flex flex-col items-center">
    <hero />
    <how />
    <badges />
    <available-on />
    <opensource-and-free />
    <integrate />
  </div>
</template>
<script>
import Hero from "~/components/Hero.vue";
import How from "~/components/How.vue";
import Badges from "~/components/Badges.vue";
import AvailableOn from "~/components/AvailableOn.vue";
import OpensourceAndFree from "~/components/OpensourceAndFree.vue";
import Integrate from "~/components/Integrate.vue";
export default {
  components: {
    Hero,
    How,
    Badges,
    AvailableOn,
    OpensourceAndFree,
    Integrate
  },
  data() {
    return {
      pageDescription:
        "nano address visualisation with a twist. available on natrium & nanocrawler.",
      pageTitle: "natricon | meet your nano address",
      pagePreview: "https://natricon.com/images/previews/preview-home.png",
      pageThemeColor: "#FFFFFF",
      canonicalURL: "https://natricon.com"
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
        },
        // Windows 8.1 + IE11 and above
        {
          name: "apple-mobile-web-app-status-bar-style",
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
        {
          rel: "icon",
          sizes: "32x32",
          href: "/favicon-32x32.png"
        },
        {
          rel: "icon",
          sizes: "16x16",
          href: "/favicon-16x16.png"
        },
        { rel: "manifest", href: "/site.webmanifest" }
      ]
    };
  }
};
</script>