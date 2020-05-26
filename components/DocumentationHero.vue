<template>
  <div class="w-full flex flex-col items-center pt-4 pb-8 md:pb-12 md:pt-12">
    <div class="n-container flex flex-col items-center px-5">
      <h1 class="text-4xl md:text-5xl text-center leading-none">
        <span class="relative inline-block">
          <span class="font-bold line-lightOrange">api</span>
        </span> documentation
      </h1>
      <h2
        class="text-xl md:text-2xl text-center mt-3"
      >we have a single endpoint and just a couple of parameters.</h2>
    </div>
  </div>
</template>
<style scoped>
.line-lightOrange::after {
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
  background-color: #ffc89f;
  z-index: -1;
}
</style>