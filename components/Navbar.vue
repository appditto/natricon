<template>
  <div class="w-full flex flex-row justify-center items-center px-4 py-4">
    <div class="n-container flex flex-row justify-between items-center">
      <div>
        <img
          class="w-36 h-auto"
          :src="require('~/assets/images/logos/natricon-logo.svg')"
          alt="Natricon Logo"
        />
      </div>
      <div class="flex flex-row font-medium">
        <a class="text-lg" href="https://github.com/appditto/natricon" target="_blank">github</a>
        <nuxt-link to="/documentation" class="text-lg ml-6 mr-1">api</nuxt-link>
      </div>
    </div>
  </div>
</template>