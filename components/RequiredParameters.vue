<template>
  <div class="w-full flex flex-col items-center pt-16 pb-8 md:pb-16">
    <div class="n-container flex flex-col items-center px-5">
      <h1 class="text-4xl md:text-5xl text-center leading-none">
        <span class="relative inline-block">
          <span class="font-bold line-brightPink">required</span>
        </span> parameters
      </h1>
      <div class="flex flex-row flex-wrap items-center mt-10 px-2">
        <code class="bg-brightPink px-3 py-1 text-xl font-bold rounded-lg">address</code>
        <span class="text-2xl font-bold m-3">:</span>
        <div class="flex flex-col">
          <h6 class="text-lg">nano address. example:</h6>
          <pre
            class="max-w-sm bg-black text-brightPink px-3 py-2 font-bold rounded-lg mt-2 break-all"
          >nano_1natrium1o3z5519ifou7xii8crpxpk8y65qmkih8e8bpsjri651oza8imdd</pre>
        </div>
      </div>
      <!-- Arrow Down -->
      <div class="w-full flex flex-row justify-center items-center mt-6 mb-2">
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
      <div class="w-full flex flex-row flex-wrap justify-center items-center mt-6">
        <code class="bg-aquaGreen px-3 py-2 text-xl font-bold rounded-lg m-4">get</code>
        <pre
          class="max-w-md md:w-120 text-xl bg-black text-aquaGreen px-4 py-2 font-bold rounded-lg mt-1 break-all"
        >https://natricon.com/api/v1/nano<span class="text-white">?</span><span
  class="text-brightPink"
>address=nano_1natrium1o3z5519ifou7xii8crpxpk8y65qmkih8e8bpsjri651oza8imdd</span></pre>
      </div>
      <!-- Arrow Down -->
      <div class="w-full flex flex-row justify-center items-center mt-6 mb-2">
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
      <!-- Result Code & Illustration -->
      <div class="w-full flex flex-row flex-wrap justify-center items-center mt-6 mb-4">
        <pre
          class="max-w-lg w-full h-56 text-sm bg-black text-aquaGreen px-4 py-2 rounded-lg mt-1 break-all overflow-y-scroll"
        >{{natriconSvgCode}}</pre>
        <div class="w-72 h-auto mx-4 my-8 md:m-8">
          <img
            class="w-full h-auto"
            :src="require('~/assets/images/illustrations/call-result-svg-desktop.svg')"
            alt="Natrium Natricon Svg"
          />
        </div>
      </div>
    </div>
  </div>
</template>
<style scoped>
.line-brightPink::after {
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
  background-color: #ffa4f6;
  z-index: -1;
}
</style>
<script>
export default {
  data() {
    return {
      natriconSvgCode:
        '<svg viewBox="0 0 512 512" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink"><g id="body"><path d="M124.917 184.007C124.131 151.898 152.011 128 184.13 128H327.869c32.119.0 60 23.898 59.213 56.007C384.12 304.971 317.02 384 256 384c-61.021.0-128.12-79.029-131.083-199.993z" fill="#a3cdff"/><path fill-rule="evenodd" clip-rule="evenodd" d="M179.534 343.157C202.43 369.404 229.777 384 256 384c61.021.0 128.12-79.029 131.083-199.993C387.465 168.392 381.067 154.718 370.699 144.791 371.582 148.024 372.079 151.428 372.138 154.966 374.254 283.819 304.351 369 241 369 220.297 369 198.894 359.903 179.534 343.157z" fill="#000" fill-opacity=".145531"/></g><g id="hair"><path d="M362 170c14.999 7 18 52 22 52C388.001 222 4e2 166 397 152s-15-15.5-25.5-26-6.5-20-18-26.5S325 99 313 95s-9.5-15.5-25-17c-15.499-1.5-27 2.5-40 4.5S219.5 75 207.5 79s-12 15-21.5 19.5-20.999 2-31.999 8S146.5 123 139 132s-19.331 14-23.999 28C110.333 174 123.999 222 128 222s12.501-49 22-52c9.5-3 20 12 30.5 9s10.001-16 21.5-17.5S227.001 182 239.5 179s29.501-18.5 40.5-18.5S305 179 317.499 180c14.736 1.179 29.502-17 44.501-10z" fill="#002a65"/><path fill-rule="evenodd" clip-rule="evenodd" d="M384.792 190.016C383.127 181.364 379.262 161.286 359 157 343.529 153.727 334.514 158.307 326.646 162.305 321.291 165.026 316.467 167.477 310.499 167 304.229 166.498 297.581 161.593 291.124 156.828 284.709 152.095 278.482 147.5 273 147.5 266.967 147.5 259.129 152.162 251.219 156.868 244.706 160.742 238.144 164.645 232.5 166 225.942 167.574 219.108 162.266 212.433 157.081 206.384 152.384 200.466 147.787 195 148.5 188.941 149.29 186.214 153.273 183.453 157.303 180.974 160.923 178.468 164.581 173.5 166 168.163 167.525 162.826 164.4 157.621 161.351 152.584 158.402 147.671 155.525 143 157 137.35 158.784 132.053 176.842 127.847 191.183 125.357 199.672 123.249 206.859 121.676 208.599 124.168 216.462 126.642 222 128.001 222 129.622 222 131.982 213.954 134.848 204.183 139.055 189.842 144.351 171.784 150.001 170 154.673 168.525 159.586 171.402 164.622 174.351 169.828 177.4 175.164 180.525 180.501 179 185.469 177.581 187.975 173.923 190.454 170.303 193.215 166.273 195.943 162.29 202.001 161.5 207.467 160.787 213.386 165.384 219.434 170.081 226.11 175.266 232.943 180.574 239.501 179 245.145 177.645 251.707 173.742 258.22 169.868 266.131 165.162 273.969 160.5 280.001 160.5 285.483 160.5 291.71 165.095 298.125 169.828 304.582 174.593 311.23 179.498 317.5 180 323.886 180.511 330.278 177.387 336.691 174.251 345.078 170.152 353.502 166.033 362.001 170 371.7 174.526 376.382 194.941 379.561 208.804 381.299 216.38 382.587 222 384.001 222c4.00100000000003.0 16-56 13-70C396.15 148.029 394.575 145.064 392.515 142.613 393.969 158.971 387.848 194 386 194 385.559 194 385.263 192.466 384.792 190.016z" fill="#000" fill-opacity=".344262"/></g><g id="mouth"><path fill-rule="evenodd" clip-rule="evenodd" d="M256 331c21 0 33-15 33-27s-15.281-5-33-5-33-7-33 5 12 27 33 27z" fill="#000" fill-opacity=".265106"/><path fill-rule="evenodd" clip-rule="evenodd" d="M283.44 300.72 283.194 303.675C283.071 305.144 282.283 306.472 280.956 307.113 277.365 308.848 269.047 312 256 312 243.044 312 234.751 308.892 231.121 307.15 229.745 306.49 228.926 305.114 228.8 303.594L228.56 300.72C230.513 299.653 234.216 300.304 239.193 301.179 243.821 301.993 249.551 303 256 303S268.179 301.993 272.807 301.179C277.784 300.304 281.487 299.653 283.44 300.72z" fill="#fff"/><path fill-rule="evenodd" clip-rule="evenodd" d="M280.192 317.364C275.281 323.253 266.846 327 256 327 253.93 327 251.948 326.864 250.062 326.602 250.878 320.67 259.717 316 270.5 316 274.006 316 277.306 316.494 280.192 317.364z" fill="#ff5858"/></g><g id="eye"><path d="M190.628 242H321.37C325.551 242 328.451 246.17 326.995 250.089L319.905 269.178C318.158 273.88 313.671 277 308.655 277H277.343C272.327 277 267.841 273.88 266.094 269.178L258.971 250H253.028L245.905 269.178C244.158 273.88 239.671 277 234.655 277H203.343C198.327 277 193.841 273.88 192.094 269.178l-7.09-19.089C183.548 246.17 186.447 242 190.628 242z" fill="#000"/><path d="M322.79 249.922l-48.335 22.539C275.359 272.811 276.335 273 277.343 273h31.312C311.999 273 314.99 270.92 316.155 267.785L322.79 249.922z" fill="#fff"/><path d="M242.156 267.785l6.553-17.641-48.033 22.398C201.517 272.84 202.417 273 203.345 273h31.312C238.001 273 240.992 270.92 242.156 267.785z" fill="#fff"/></g><g id="badge"><circle cx="357" cy="357" r="43" fill="#fff"/><circle cx="357" cy="357" r="36" fill="#1a82ff"/><path d="M372.333 365.123C376.559 365.123 380 361.71 380 357.467S376.582 349.811 372.333 349.811C366.583 349.811 364.667 347.897 364.667 342.156 364.667 337.936 361.226 334.5 357 334.5S349.333 337.913 349.333 342.156C349.333 347.897 347.417 349.811 341.667 349.811 337.441 349.811 334 353.224 334 357.467 334 361.687 337.441 365.123 341.667 365.123S349.333 361.71 349.333 357.467C349.333 351.725 351.25 349.811 357 349.811S364.667 351.725 364.667 357.467C364.667 361.687 368.107 365.123 372.333 365.123z" fill="#fefefe"/><path d="M357.001 379.592C361.235 379.592 364.668 376.164 364.668 371.936 364.668 367.708 361.235 364.28 357.001 364.28 352.767 364.28 349.335 367.708 349.335 371.936 349.335 376.164 352.767 379.592 357.001 379.592z" fill="#fefefe"/></g></svg>'
    };
  }
};
</script>