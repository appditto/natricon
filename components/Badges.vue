<template>
  <div class="w-full flex flex-col items-center pt-16 pb-20" id="badges">
    <div class="n-container flex flex-col items-center px-5">
      <h3 class="text-4xl md:text-5xl text-center leading-none">
        exchanges? services? nodes?
        <br />we have a
        <span class="relative inline-block">
          <span class="font-bold line-green">badge</span>
        </span> for it.
      </h3>
      <h4 class="text-xl md:text-2xl text-center mt-3">
        so, you can identify the category
        <br class="hidden md:block" />of an address at a glance.
      </h4>
      <div
        class="w-full max-w-md md:max-w-full flex flex-row flex-wrap mt-4 md:mt-8 md:px-8 lg:px-20"
      >
        <div class="w-1/2 md:w-1/4 p-3">
          <img
            class="w-full h-auto"
            :src="require('~/assets/images/illustrations/badge-exchange.svg')"
            alt="Exchange Badge"
          />
        </div>
        <div class="w-1/2 md:w-1/4 p-3">
          <img
            class="w-full h-auto"
            :src="require('~/assets/images/illustrations/badge-service.svg')"
            alt="Service Badge"
          />
        </div>
        <div class="w-1/2 md:w-1/4 p-3">
          <img
            class="w-full h-auto"
            :src="require('~/assets/images/illustrations/badge-node.svg')"
            alt="Node Badge"
          />
        </div>
        <div class="w-1/2 md:w-1/4 p-3">
          <img
            class="w-full h-auto"
            :src="require('~/assets/images/illustrations/badge-donor.svg')"
            alt="Donor Badge"
          />
        </div>
      </div>
    </div>
  </div>
</template>
<style scoped>
.line-green::after {
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
  background-color: #66ffb2;
  z-index: -1;
}
</style>