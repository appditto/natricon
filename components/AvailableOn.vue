<template>
  <div class="w-full flex flex-col items-center py-8 md:py-20 relative">
    <!-- Right Illustration -->
    <img
      class="hidden md:block bg-illustration h-auto absolute right-0"
      :src="require('~/assets/images/illustrations/natrium-and-nanocrawler-right.svg')"
      alt="Background Illustration Right"
    />
    <!-- Left Illustration -->
    <img
      class="hidden md:block bg-illustration h-auto absolute left-0 bottom-0"
      :src="require('~/assets/images/illustrations/natrium-and-nanocrawler-left.svg')"
      alt="Background Illustration Left"
    />
    <!-- Main Part -->
    <div class="n-container flex flex-col items-center px-5 z-10">
      <h3 class="text-4xl md:text-5xl text-center leading-tight">
        available on
        <br />
        <span class="relative inline-block">
          <span class="font-bold line-lightBlue">natrium</span>
        </span> &
        <span class="relative inline-block">
          <span class="font-bold line-green">nanocrawler</span>
        </span>
      </h3>
      <h4 class="text-xl md:text-2xl text-center mt-3">coming soon to more nano services.</h4>
      <div class="flex flex-row flex-wrap justify-center mt-2">
        <a class="w-full md:w-72 mt-5 mx-3" href="https://natrium.io" target="_blank">
          <button
            class="w-full bg-black text-white hover:text-lightBlue btn-shadow-lightBlue font-medium text-xl rounded-full px-6 pt-1 pb-3"
          >download natrium</button>
        </a>
        <a class="w-full md:w-72 mt-5 mx-3" href="https://nanocrawler.cc" target="_blank">
          <button
            class="w-full bg-black text-white hover:text-green btn-shadow-green font-medium text-xl rounded-full pt-1 pb-3"
          >visit nanocrawler</button>
        </a>
      </div>
      <div class="w-full flex flex-col items-center mt-12 md:mt-0">
        <div class="w-full max-w-md md:max-w-full px-16 md:hidden">
          <img
            class="w-full h-auto"
            :src="require('~/assets/images/illustrations/natrium.svg')"
            alt="Natrium"
          />
        </div>
        <div class="w-full max-w-md md:max-w-full pl-2 pr-12 md:hidden mt-10">
          <img
            class="w-full h-auto"
            :src="require('~/assets/images/illustrations/nanocrawler-account.svg')"
            alt="Nanocrawler Account"
          />
        </div>
        <div class="w-full max-w-md md:max-w-full pl-12 pr-2 md:hidden mt-6">
          <img
            class="w-full h-auto"
            :src="require('~/assets/images/illustrations/nanocrawler-send.svg')"
            alt="Nanocrawler Send"
          />
        </div>
      </div>
      <div class="w-full md:px-16 lg:px-24 mt-16 hidden md:block">
        <img
          class="w-full h-auto"
          :src="require('~/assets/images/illustrations/natrium-and-nanocrawler-desktop.svg')"
          alt="Natrium and Nanocrawler"
        />
      </div>
    </div>
  </div>
</template>
<style scoped>
.btn-shadow-lightBlue {
  box-shadow: -0.3rem 0.4rem 0rem 0rem#a3cdff;
  transition: all 0.2s ease-out;
  transform: scale(1);
}
.btn-shadow-green {
  box-shadow: -0.3rem 0.4rem 0rem 0rem#66ffb2;
  transition: all 0.2s ease-out;
  transform: scale(1);
}
.btn-shadow-lightBlue:hover {
  box-shadow: 0rem 0rem 0rem 0rem#a3cdff;
  transform: scale(0.95);
}
.btn-shadow-green:hover {
  box-shadow: 0rem 0rem 0rem 0rem#66ffb2;
  transform: scale(0.95);
}
.line-lightBlue::before {
  display: block;
  position: absolute;
  width: calc(100% + 0.3rem);
  left: -0.15rem;
  content: "";
  height: 0.75rem;
  border-radius: 0.15rem;
  top: 100%;
  margin-left: auto;
  margin-right: auto;
  margin-top: -0.9rem;
  background-color: #a3cdff;
  z-index: -1;
}
.line-green::before {
  display: block;
  position: absolute;
  width: calc(100% + 0.15rem);
  left: -0.15rem;
  content: "";
  height: 0.75rem;
  border-radius: 0.15rem;
  top: 100%;
  margin-left: auto;
  margin-right: auto;
  margin-top: -0.9rem;
  background-color: #66ffb2;
  z-index: -1;
}
.bg-illustration {
  max-width: calc((100vw - 1280px) / 2.5);
  min-width: 12rem;
  width: 100%;
}
@media only screen and (min-width: 1024px) {
  .bg-illustration {
    min-width: 16rem;
  }
}
</style>