<template>
  <div class="flex flex-col shadow-xl rounded-lg">
    <div class="w-56 h-56">
      <div class="w-full relative w-56 h-56">
        <div v-html="svg"></div>
      </div>
    </div>
    <!-- Color Info -->
    <div class="w-64 flex flex-row mt-5 justify-around font-mono">
      <div class="w-1/4 flex flex-col items-end">
        <div class="flex flex-col items-start pr-1 py-1">
          <span class="font-bold">&nbsp;</span>
          <span class="font-bold">Body:</span>
          <span class="font-bold">Hair:</span>
          <span class="font-bold">Delta:</span>
        </div>
      </div>
      <div class="w-1/4 flex flex-col items-end px-1">
        <div class="w-full flex flex-col items-end bg-black-05 rounded px-2 py-1">
          <span class="font-bold">H</span>
          <span>{{bodyH}}</span>
          <span>{{hairH}}</span>
          <span>{{deltaH}}</span>
        </div>
      </div>
      <div class="w-1/4 flex flex-col items-end px-1">
        <div class="w-full flex flex-col items-end bg-black-05 rounded px-2 py-1">
          <span class="font-bold">S</span>
          <span>{{bodyS}}</span>
          <span>{{hairS}}</span>
          <span>{{deltaS}}</span>
        </div>
      </div>
      <div class="w-1/4 flex flex-col items-end px-1">
        <div class="w-full flex flex-col items-end bg-black-05 rounded px-2 py-1">
          <span class="font-bold">B</span>
          <span>{{bodyV}}</span>
          <span>{{hairV}}</span>
          <span>{{deltaV}}</span>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  props: {
    svg: null,
    bodyH: null,
    bodyS: null,
    bodyV: null,
    deltaH: null,
    deltaS: null,
    deltaV: null,
    hairH: null,
    hairS: null,
    hairV: null
  }
};
</script>
<style>
.bg-black-05 {
  background-color: rgba(0, 0, 0, 0.05);
}
</style>
