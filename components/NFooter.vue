<template>
  <div
    class="w-full flex flex-row justify-center items-center px-4 py-8 md:px-6"
    id="footer"
  >
    <div class="n-container flex flex-col justify-center items-center">
      <div class="flex flex-row flex-wrap justify-center">
        <a
          class="relative m-3"
          href="https://twitter.com/appditto"
          target="_blank"
        >
          <div class="line bga-yellow scaleY w-12 h-12">
            <img
              class="w-full h-full"
              :src="require('~/assets/images/icons/twitter.svg')"
              alt="Twitter Icon"
            />
          </div>
        </a>
        <a
          class="relative m-3"
          href="https://medium.com/appditto"
          target="_blank"
        >
          <div class="w-12 h-12 line bga-yellow scaleY">
            <img
              class="w-full h-full"
              :src="require('~/assets/images/icons/medium.svg')"
              alt="Medium Icon"
            />
          </div>
        </a>
        <a
          class="relative m-3"
          href="https://instagram.com/appditto"
          target="_blank"
        >
          <div class="w-12 h-12 line bga-yellow scaleY">
            <img
              class="w-full h-full"
              :src="require('~/assets/images/icons/instagram.svg')"
              alt="Instagram Icon"
            />
          </div>
        </a>
        <a
          class="relative m-3"
          href="https://github.com/appditto"
          target="_blank"
        >
          <div class="w-12 h-12 line bga-yellow scaleY">
            <img
              class="w-full h-full"
              :src="require('~/assets/images/icons/github.svg')"
              alt="GitHub Icon"
            />
          </div>
        </a>
      </div>
      <h6 class="font-medium mt-2 md:mt-3 mb-3 md:mb-4">©2022, natricon.</h6>
      <div class="flex flex-row justify-center my-2 md:my-3">
        <span class="font-medium mx-1 mt-1">made by</span>
        <a class="relative mx-1" href="https://appditto.com" target="_blank">
          <div class="w-32 h-auto line-2 bga-green scaleY">
            <img
              class="w-full h-full"
              :src="require('~/assets/images/logos/appditto-logo.svg')"
              alt="Appditto Logo"
            />
          </div>
        </a>
      </div>
    </div>
  </div>
</template>
<style scoped>
.bga-green::after {
  background-color: #66ffb2;
}
.bga-yellow::after {
  background-color: #ffee52;
}
.line::after {
  display: block;
  position: absolute;
  width: calc(100% + 0.4rem);
  left: -0.2rem;
  content: '';
  height: 1.75rem;
  border-radius: 0.15rem;
  margin-left: auto;
  margin-right: auto;
  margin-top: -1.6rem;
  z-index: -1;
  transition: all 0.2s ease-out;
  transform-origin: center bottom;
  transform: scaleY(0);
}
.line-2::after {
  display: block;
  position: absolute;
  width: calc(100% + 0.4rem);
  left: -0.2rem;
  content: '';
  height: 1rem;
  border-radius: 0.15rem;
  margin-left: auto;
  margin-right: auto;
  margin-top: -0.7rem;
  z-index: -1;
  transition: all 0.2s ease-out;
  transform-origin: center bottom;
  transform: scaleY(0);
}
.scaleY:hover::after {
  transform: scaleY(1);
}
</style>