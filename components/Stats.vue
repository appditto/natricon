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
              class="w-full flex flex-col flex-1 justify-center items-center bg-green rounded-lg pt-6 pb-10 px-4 md:px-8 border-4 border-black card"
            >
              <h6 class="text text-center">we served</h6>
              <h5 class="text-5xl font-bold text-center leading-tight">{{ uniqueServed }}</h5>
              <h6 class="text text-center mt-2">unique natricons</h6>
            </div>
          </div>
          <div class="w-full md:w-1/2 lg:w-1/3 max-w-sm px-4 mt-8 flex flex-col">
            <div
              class="w-full flex flex-col flex-1 justify-center items-center bg-green rounded-lg pt-6 pb-10 px-4 md:px-8 border-4 border-black card"
            >
              <div @mouseleave="isTooltipVisible=false" class="flex flex-col">
                <h6 class="text text-center">available on</h6>
                <div class="relative w-48">
                  <transition name="tooltipTransition">
                    <div
                      v-if="isTooltipVisible"
                      class="w-48 z-50 absolute flex flex-col bg-black p-2 rounded-lg text-white tooltip transform bottom-0 mb-16"
                    >
                      <p
                        class="p-2 text-green font-bold text-lg hover:bg-green hover:text-black rounded-sm transition-colors duration-150 ease-out"
                      >natricon.com</p>
                      <a
                        href="https://natrium.io"
                        target="_blank"
                        class="p-2 text-green font-bold text-lg hover:bg-green hover:text-black rounded-sm transition-colors duration-150 ease-out"
                      >natrium</a>
                      <a
                        href="https://nanocrawler.cc"
                        target="_blank"
                        class="p-2 text-green font-bold text-lg hover:bg-green hover:text-black rounded-sm transition-colors duration-150 ease-out"
                      >nanocrawler</a>
                      <a
                        href="https://plinko.somenano.com"
                        target="_blank"
                        class="p-2 text-green font-bold text-lg hover:bg-green hover:text-black rounded-sm transition-colors duration-150 ease-out"
                      >plinko</a>
                      <a
                        href="https://sendr.space"
                        target="_blank"
                        class="p-2 text-green font-bold text-lg hover:bg-green hover:text-black rounded-sm transition-colors duration-150 ease-out"
                      >sendr.space</a>
                    </div>
                  </transition>
                  <p
                    @mouseover="isTooltipVisible=true"
                    class="text-5xl font-bold text-center leading-tight cursor-pointer"
                  >{{ serviceCount }}</p>
                </div>
              </div>

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
  },
  data() {
    return {
      isTooltipVisible: false
    };
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
.left-half {
  left: 50%;
}
.tooltip::after {
  content: " ";
  position: absolute;
  top: 100%; /* At the top of the tooltip */
  left: 50%;
  margin-left: -10px;
  border-width: 10px;
  border-style: solid;
  border-color: black transparent transparent transparent;
}
.tooltipTransition-enter-active {
  transition: all 0.2s ease-out;
  transform-origin: bottom center;
}
.tooltipTransition-leave-active {
  transition: all 0.2s ease-out;
  transform-origin: bottom center;
}
.tooltipTransition-enter {
  opacity: 0;
  transform: translateY(1rem) scale(0.75);
}
.tooltipTransition-leave-to {
  opacity: 0;
  transform: scale(0.75);
}
</style>