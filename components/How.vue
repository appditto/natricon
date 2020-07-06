<template>
  <div class="w-full flex flex-col items-center relative pt-4 pb-4" id="how">
    <!-- Right Illustration -->
    <img
      class="hidden md:block bg-illustration h-auto absolute right-0 opacity-15 md:opacity-100"
      :src="require('~/assets/images/illustrations/natricon-formula-right.svg')"
      alt="Background Illustration Right"
    />
    <!-- Left Illustration -->
    <img
      class="hidden md:block bg-illustration h-auto absolute left-0 bottom-0 opacity-15 md:opacity-100"
      :src="require('~/assets/images/illustrations/natricon-formula-left.svg')"
      alt="Background Illustration Left"
    />
    <!-- Main Part -->
    <div
      class="n-container flex flex-col items-center px-5 z-10 my-8 md:mt-24 md:mb-40 lg:mt-32 lg:mb-48"
    >
      <h3 class="text-4xl md:text-5xl text-center leading-none">
        how does it
        <br class="md:hidden" />
        <span class="relative inline-block">
          <span class="font-bold line-redPink">work?</span>
        </span>
      </h3>
      <h4 class="text-xl md:text-2xl text-center mt-3">
        we take your nano address, apply some math,
        <br class="hidden md:block" />and it magically turns into
        <span class="font-bold">natricon</span>.
      </h4>
      <div class="w-full max-w-md md:max-w-full px-16 md:px-24 md:hidden mt-8">
        <img
          class="w-full h-auto"
          :src="require('~/assets/images/illustrations/natricon-formula-mobile.svg')"
          alt="Natricon Formula Mobile"
        />
      </div>
      <div class="w-full md:px-16 lg:px-24 mt-12 hidden md:block">
        <img
          class="w-full h-auto"
          :src="require('~/assets/images/illustrations/natricon-formula-desktop.svg')"
          alt="Natricon Formula Desktop"
        />
      </div>
      <!-- Youtube Video -->
      <div class="w-full md:max-w-xl lg:max-w-4xl mt-16 md:mt-20 lg:mt-24">
        <lazy-youtube-video
          :iframeAttributes="{title: 'Introduction Video'}"
          class="w-full h-auto rounded-lg overflow-hidden border-3 lg:border-4 border-black"
          alt="Introduction Video"
          src="https://www.youtube.com/embed/ZtaTEqTtukk"
        />
      </div>
    </div>
  </div>
</template>
<script>
import LazyYoutubeVideo from "vue-lazy-youtube-video";
import "vue-lazy-youtube-video/dist/style.css";
export default {
  components: {
    LazyYoutubeVideo
  }
};
</script>
<style scoped>
.bg-shadow {
  box-shadow: 0rem 0rem 0rem 0.2rem#000000;
}
@media only screen and (min-width: 768px) {
  .bg-shadow {
    box-shadow: 0rem 0rem 0rem 0.3rem#000000;
  }
}
.line-redPink::after {
  display: block;
  position: absolute;
  width: calc(100% + 0.3rem);
  left: -0.15rem;
  content: "";
  height: 0.75rem;
  border-radius: 0.15rem;
  margin-left: auto;
  margin-right: auto;
  margin-top: -0.6rem;
  background-color: #fe98cd;
  z-index: -1;
}
.bg-illustration {
  max-width: calc((100vw - 1280px) / 2.75);
  min-width: 12rem;
  width: 100%;
}
@media only screen and (min-width: 768px) {
  .bg-illustration {
    min-width: 12rem;
  }
}
@media only screen and (min-width: 1024px) {
  .bg-illustration {
    min-width: 16rem;
  }
}
</style>