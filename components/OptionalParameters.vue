<template>
  <div class="w-full flex flex-col items-center py-16">
    <div class="n-container flex flex-col items-center px-5">
      <h1 class="text-4xl md:text-5xl text-center leading-none mb-6">
        <span class="relative inline-block">
          <span class="font-bold line-lime">optional</span>
        </span> parameters
      </h1>
      <!-- Format -->
      <div class="w-full flex flex-row flex-wrap justify-center items-center my-6 px-2">
        <div class="flex flex-row w-full md:w-1/3 md:justify-end items-center">
          <code class="bg-lime px-3 py-1 text-xl font-bold rounded-lg my-3">format</code>
          <span class="text-2xl font-bold mx-3">:</span>
        </div>
        <div class="flex flex-row w-full md:w-1/2">
          <span class="text-lg">
            <code class="font-bold bg-black text-lime px-1_5 py-0_5 rounded-md">svg</code> (default).
            <br />supported parameters are
            <code
              class="font-bold bg-black text-lime px-1_5 py-0_5 rounded-md"
            >png</code>,
            <code class="font-bold bg-black text-lime px-1_5 py-0_5 rounded-md">webp</code> &
            <code class="font-bold bg-black text-lime px-1_5 py-0_5 rounded-md">svg</code>.
          </span>
        </div>
      </div>
      <!-- Size -->
      <div class="w-full flex flex-row flex-wrap justify-center items-center my-6 px-2">
        <div class="flex flex-row w-full md:w-1/3 md:justify-end items-center">
          <code class="bg-lime px-3 py-1 text-xl font-bold rounded-lg my-3">size</code>
          <span class="text-2xl font-bold mx-3">:</span>
        </div>
        <div class="flex flex-row w-full md:w-1/2">
          <span class="text-lg leading-loose">
            in pixels. ignored when format is
            <code
              class="font-bold bg-black text-lime px-1_5 py-0_5 rounded-md"
            >svg</code>.
            <br />defaults to
            <code class="font-bold bg-black text-lime px-1_5 py-0_5 rounded-md">128</code> when format is
            <code class="font-bold bg-black text-lime px-1_5 py-0_5 rounded-md">webp</code> or
            <code class="font-bold bg-black text-lime px-1_5 py-0_5 rounded-md">png</code>.
            <br />minimum is
            <code class="font-bold bg-black text-lime px-1_5 py-0_5 rounded-md">100</code>, maximum is
            <code class="font-bold bg-black text-lime px-1_5 py-0_5 rounded-md">1000</code>.
          </span>
        </div>
      </div>
      <!-- Outline -->
      <div class="w-full flex flex-row flex-wrap justify-center items-center my-6 px-2">
        <div class="flex flex-row w-full md:w-1/3 md:justify-end items-center">
          <code class="bg-lime px-3 py-1 text-xl font-bold rounded-lg my-3">outline</code>
          <span class="text-2xl font-bold mx-3">:</span>
        </div>
        <div class="flex flex-row w-full md:w-1/2">
          <span class="text-lg leading-loose">
            <code class="font-bold bg-black text-lime px-1_5 py-0_5 rounded-md">false</code>
            (default).
            <br />when
            <code class="font-bold bg-black text-lime px-1_5 py-0_5 rounded-md">true</code>, a
            <span class="font-bold">white</span> outline will be added to the natricon.
          </span>
        </div>
      </div>
      <!-- Outline Color -->
      <div class="w-full flex flex-row flex-wrap justify-center items-center my-6 px-2">
        <div class="flex flex-row w-full md:w-1/3 md:justify-end items-center">
          <code class="bg-lime px-3 py-1 text-xl font-bold rounded-lg my-3">outline_color</code>
          <span class="text-2xl font-bold mx-3">:</span>
        </div>
        <div class="flex flex-row w-full md:w-1/2">
          <span class="text-lg leading-loose">
            <code class="font-bold bg-black text-lime px-1_5 py-0_5 rounded-md">white</code> (default).
            <br />supported parameters are
            <code
              class="font-bold bg-black text-lime px-1_5 py-0_5 rounded-md"
            >white</code> &
            <code class="font-bold bg-black text-lime px-1_5 py-0_5 rounded-md">black</code>.
            <br />Used when
            <code
              class="font-bold bg-black text-lime px-1_5 py-0_5 rounded-md"
            >outline=true</code>.
          </span>
        </div>
      </div>
      <!-- Arrow Down -->
      <div class="w-full flex flex-row flex-wrap justify-center items-center mt-2 mb-6">
        <img
          class="w-12 h-12"
          :src="require('~/assets/images/icons/arrow-down.svg')"
          alt="Arrow Down"
        />
      </div>
      <!-- Example Call -->
      <h2 class="text-3xl md:text-4xl text-center leading-none mt-2">
        <b>example</b> call
      </h2>
      <div class="w-full flex flex-row flex-wrap flex-wrap justify-center items-center mt-6">
        <code class="bg-aquaGreen px-3 py-2 text-xl font-bold rounded-lg my-3 mx-4">get</code>
        <pre
          class="w-128 text-xl bg-black text-lime px-4 py-2 font-bold rounded-lg mt-1 break-all my-2"
        ><span class="text-aquaGreen">https://natricon.com/api/v1/nano</span><span class="text-white">?</span><span
  class="text-brightPink"
>address=nano_1natrium1o3z5519ifou7xii8crpxpk8y65qmkih8e8bpsjri651oza8imdd</span><span
  class="text-white"
>&</span>format=png<span class="text-white">&</span>size=512<span class="text-white">&</span>outline=true<span class="text-white">&</span>outline_color=black</pre>
      </div>
      <!-- Arrow Down -->
      <div class="w-full flex flex-row flex-wrap justify-center items-center mt-6 mb-4">
        <img
          class="w-12 h-12"
          :src="require('~/assets/images/icons/arrow-down.svg')"
          alt="Arrow Down"
        />
      </div>
      <!-- Example Result -->
      <h2 class="text-3xl md:text-4xl text-center leading-none mt-2">
        <b>example</b> result
      </h2>
      <!-- Illustration -->
      <div
        class="w-full flex flex-row flex-wrap justify-center items-center mt-8 mb-4 md:px-16 lg:px-48"
      >
        <img
          class="max-w-xs w-full h-auto lg:hidden"
          :src="require('~/assets/images/illustrations/call-result-png-mobile.svg')"
          alt="Natrium Natricon Png Mobile"
        />
        <img
          class="w-full h-auto hidden lg:block"
          :src="require('~/assets/images/illustrations/call-result-png-desktop.svg')"
          alt="Natrium Natricon Png Desktop"
        />
      </div>
    </div>
  </div>
</template>
<script>
import Divider from "~/components/Divider.vue";
export default {
  components: {
    Divider
  }
};
</script>
<style scoped>
.line-lime::after {
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
  background-color: #d4ff00;
  z-index: -1;
}
</style>