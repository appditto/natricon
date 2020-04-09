<template>
  <div class="flex flex-col shadow-xl rounded-lg">
    <div class="w-56 h-56">
      <div class="w-full relative w-56 h-56">
        <!-- Body -->
        <svg class="absolute" viewBox="0 0 512 512" fill="none" xmlns="http://www.w3.org/2000/svg">
          <rect x="128" y="128" width="256" height="256" rx="48" :fill="bodyColor" />
          <path
            fill-rule="evenodd"
            clip-rule="evenodd"
            d="M138.883 366.438C147.686 377.16 161.044 384 176 384H336C362.51 384 384 362.51 384 336V176C384 161.044 377.16 147.686 366.438 138.883C368.094 143.273 369 148.031 369 153V329C369 351.091 351.091 369 329 369H153C148.031 369 143.273 368.094 138.883 366.438Z"
            fill="black"
            fill-opacity="0.15"
          />
        </svg>
        <!-- Hair -->
        <svg class="absolute" viewBox="0 0 512 512" fill="none" xmlns="http://www.w3.org/2000/svg">
          <path
            d="M302 194C359 194 408 177 408 132.5C408 50.5001 317 98.6221 223.5 92.5C158.76 88.2611 120.816 111.709 120 161.5C119.5 192 129 225.5 129 225.5C129 225.5 124 189.5 149 178C174 166.5 245 194 302 194Z"
            :fill="hairColor"
          />
          <path
            fill-rule="evenodd"
            clip-rule="evenodd"
            d="M381 122.5C381 167 354 179 297 179C272.188 179 244.724 174.547 219.556 170.467C186.905 165.173 158.118 160.506 144 167C128.808 173.989 124.694 190.394 123.805 202.191C126.263 215.849 129 225.5 129 225.5C129 225.5 124 189.5 149 178C162.691 171.702 190.179 177.101 221.608 183.274C247.568 188.373 276.216 194 302 194C359 194 408 177 408 132.5C408 100.229 393.906 88.1117 371.113 84.9097C377.537 92.2917 381 104.131 381 122.5Z"
            fill="black"
            fill-opacity="0.15"
          />
        </svg>
        <!-- Mouth -->
        <svg class="absolute" viewBox="0 0 512 512" fill="none" xmlns="http://www.w3.org/2000/svg">
          <path
            fill-rule="evenodd"
            clip-rule="evenodd"
            d="M256 326C270.359 326 282 316.15 282 304C266.135 304 245.846 304 230 304C230 316.15 241.641 326 256 326Z"
            fill="black"
            fill-opacity="0.3"
          />
        </svg>
        <!-- Eyes -->
        <svg class="absolute" viewBox="0 0 512 512" fill="none" xmlns="http://www.w3.org/2000/svg">
          <path
            d="M221 273C229.837 273 237 265.837 237 257C237 248.163 229.837 241 221 241C212.163 241 205 248.163 205 257C205 265.837 212.163 273 221 273Z"
            fill="black"
            fill-opacity="0.3"
          />
          <path
            d="M291 273C299.837 273 307 265.837 307 257C307 248.163 299.837 241 291 241C282.163 241 275 248.163 275 257C275 265.837 282.163 273 291 273Z"
            fill="black"
            fill-opacity="0.3"
          />
        </svg>
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
    bodyColor: null,
    hairColor: null,
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
