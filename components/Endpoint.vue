<template>
  <div class="w-full flex flex-col items-center pt-12 pb-16">
    <div class="n-container flex flex-col items-center px-5">
      <h1 class="text-4xl md:text-5xl text-center leading-none">
        <span class="relative inline-block">
          <span class="font-bold line-aquaGreen">endpoint</span>
        </span>
      </h1>
      <div class="w-full flex flex-row flex-wrap justify-center items-center mt-6">
        <code class="bg-aquaGreen px-3 py-2 text-xl font-bold rounded-lg m-2">get</code>
        <code
          class="bg-black text-aquaGreen px-3 py-2 text-lg md:text-xl font-bold rounded-lg overflow-x-scroll my-2"
        >https://natricon.com/api/v1/nano</code>
      </div>
    </div>
  </div>
</template>
<style scoped>
.line-aquaGreen::before {
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
  margin-top: -0.6rem;
  background-color: #7bffc8;
  z-index: -1;
}
</style>