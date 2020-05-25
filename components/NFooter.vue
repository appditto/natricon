<template>
  <div class="w-full flex flex-row justify-center items-center px-4 py-4 md:px-6 md:py-6">
    <div class="n-container flex flex-col justify-center items-center">
      <div class="flex flex-row flex-wrap justify-center">
        <a href="https://twitter.com/appditto">
          <img
            class="w-12 h-12 m-3"
            :src="require('~/assets/images/icons/twitter.svg')"
            alt="Twitter Icon"
          />
        </a>
        <a href="https://medium.com/appditto">
          <img
            class="w-12 h-12 m-3"
            :src="require('~/assets/images/icons/medium.svg')"
            alt="Medium Icon"
          />
        </a>
        <a href="https://instagram.com/appditto">
          <img
            class="w-12 h-12 m-3"
            :src="require('~/assets/images/icons/instagram.svg')"
            alt="Instagram Icon"
          />
        </a>
        <a href="https://github.com/appditto">
          <img
            class="w-12 h-12 m-3"
            :src="require('~/assets/images/icons/github.svg')"
            alt="GitHub Icon"
          />
        </a>
      </div>
      <h6 class="font-medium my-2 md:my-3">©2020, natricon.</h6>
      <div class="flex flex-row justify-center my-2 md:my-3">
        <span class="font-medium mx-1 mt-1">made by</span>
        <img
          class="w-32 h-auto mx-1"
          :src="require('~/assets/images/logos/appditto-logo.svg')"
          alt="Appditto Logo"
        />
      </div>
    </div>
  </div>
</template>