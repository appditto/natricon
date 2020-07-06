<template>
  <div class="w-full flex flex-col items-center pt-10 pb-16 md:pt-16 md:pb-56 relative" id="stats">
    <!-- Main Part -->
    <div class="n-container flex flex-col items-center px-5 md:px-12 lg:px-24 z-10">
      <h3 class="text-4xl md:text-5xl text-center leading-tight">
        <span class="relative inline-block">
          <span class="font-bold line-green">natricon</span>
        </span> stats
      </h3>
      <div class="w-full relative mt-2 md:mt-4">
        <!-- Background -->
        <div
          class="w-full hidden md:block bg-white bg-shadow border-4 border-black rounded-full h-20 absolute bottom-0 z-0 -mb-10"
        ></div>
        <!-- Stats -->
        <div class="w-full flex flex-row flex-wrap justify-center z-50 relative md:px-8">
          <div class="w-full md:w-1/2 lg:w-1/3 max-w-sm px-4 mt-8 flex flex-col">
            <div
              class="w-full flex flex-col flex-1 justify-center items-center bg-green rounded-lg pt-6 pb-10 px-4 md:px-8 border-4 border-black card overflow-hidden"
            >
              <h6 class="text text-center">we served</h6>
              <h5 class="text-5xl font-bold text-center leading-tight">{{ uniqueServed }}</h5>
              <h6 class="text text-center mt-2">unique natricons</h6>
            </div>
          </div>
          <div class="w-full md:w-1/2 lg:w-1/3 max-w-sm px-4 mt-8 flex flex-col">
            <div
              class="w-full flex flex-col flex-1 justify-center items-center bg-green rounded-lg pt-6 pb-10 px-4 md:px-8 border-4 border-black card overflow-hidden"
            >
              <h6 class="text text-center">available on</h6>
              <h5 class="text-5xl font-bold text-center leading-tight">{{ serviceCount }}</h5>
              <h6 class="text text-center mt-2">services</h6>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  props: {
    serviceCount: Number,
    uniqueServed: String
  }
};
</script>
<style scoped>
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
.card {
  box-shadow: -0.4rem 0.4rem 0rem 0rem#ffffff, -0.8rem 0.8rem 0rem 0rem#000000;
}
.bg-shadow {
  box-shadow: -0.4rem 0.4rem 0rem 0rem#000000;
}
</style>