<template>
  <div class="w-full flex flex-col items-center py-8 md:py-24">
    <div class="n-container flex flex-col items-center px-5">
      <h3 class="text-4xl md:text-5xl text-center leading-tight">
        <span class="relative inline-block">
          <span class="font-bold line-cyan">open-source</span>
        </span>
        <br class="md:hidden" />&
        <span class="relative inline-block">
          <span class="font-bold line-lightPink">free</span>
        </span>
      </h3>
      <h4 class="text-xl md:text-2xl text-center mt-3">
        check out our code, see if you like it. contribute to it.
        it’s free.
        <br class="hidden md:block" />if you like it, you can donate to keep us going.
        <br class="hidden md:block" />donations of
        <b>2 nano</b> will get a
        <b>donor badge for 1 month</b>.
      </h4>
      <div class="w-full flex flex-row flex-wrap justify-center mt-2 relative">
        <a
          class="w-full md:w-56 mt-5 mx-3"
          href="https://github.com/appditto/natricon"
          target="_blank"
        >
          <button
            class="btn btn-shadow-cyan hover:text-cyan w-full bg-black text-white font-medium text-xl rounded-full px-6 pt-1 pb-3"
          >visit the repo</button>
        </a>
        <button
          @click="isDropdownOpen?closeDonateDropdown():openDonateDropdown()"
          :class="isDropdownOpen?'bg-lightPink text-black btn-shadow-black':'bg-black text-white hover:text-lightPink btn-shadow-lightPink'"
          class="btn w-full md:w-56 font-medium text-xl rounded-full px-6 pt-1 pb-3 mt-5 mx-3"
        >{{isDropdownOpen?'close':'donate'}}</button>
      </div>
      <!-- Donate Dropdow Container-->
      <div class="w-full flex flex-row justify-center bg-white relative">
        <!-- Donate Dropdown -->
        <div
          :class="isDropdownOpen?'scale-y-100 border-black shadow-lightPink':'scale-y-0 border-transparent'"
          class="w-full md:w-144 absolute flex flex-col items-center max-w-128 bg-white border-4 transition-all transform origin-top duration-300 ease-out overflow-hidden mt-5 rounded-lg px-2 md:px-12 z-50"
        >
          <!-- Go Back Button -->
          <button
            v-if="isDonationInitiated"
            class="absolute left-0 top-0 mx-3 my-2 px-3 pt-1 pb-2 font-bold"
            @click="resetDonation()"
          >go back</button>
          <div
            :class="isDropdownOpen?'opacity-100 duration-1000':'opacity-0 duration-150' "
            class="w-full flex flex-col justify-center items-center ease-out pt-8 pb-6"
          >
            <img
              v-if="donationAmount>2 && donationAmount<=10"
              class="w-32 h-32"
              :src="require('~/assets/images/gifs/NatriconDonatePhase2.gif')"
              alt="Natricon Donate 2"
            />
            <img
              v-else-if="donationAmount>10 && donationAmount<100"
              class="w-32 h-32"
              :src="require('~/assets/images/gifs/NatriconDonatePhase3.gif')"
              alt="Natricon Donate 3"
            />
            <img
              v-else-if="donationAmount>=100"
              class="w-32 h-32"
              :src="require('~/assets/images/gifs/NatriconDonatePhase4.gif')"
              alt="Natricon Donate 4"
            />
            <img
              v-else
              class="w-32 h-32"
              :src="require('~/assets/images/gifs/NatriconDonatePhase1.gif')"
              alt="Natricon Donate 1"
            />
            <div class="flex flex-col items-center">
              <!-- Donate Amount Buttons -->
              <div
                v-if="!isDonationInitiated"
                class="flex flex-row flex-wrap justify-center items-center my-4"
              >
                <button
                  @mouseover="donationAmount=2"
                  @mouseleave="customNanoAmountModel?donationAmount=customNanoAmountModel:donationAmount=2"
                  @click="initiateDonationFor(2)"
                  class="btn w-32 font-medium text-lg bg-black text-white hover:text-lightPink btn-sm-shadow-lightPink rounded-lg px-3 pt-1 pb-3 mt-5 mx-3"
                >2 nano</button>
                <button
                  @mouseover="donationAmount=10"
                  @mouseleave="customNanoAmountModel?donationAmount=customNanoAmountModel:donationAmount=2"
                  @click="initiateDonationFor(10)"
                  class="btn w-32 font-medium text-lg bg-black text-white hover:text-lightPink btn-sm-shadow-lightPink rounded-lg px-3 pt-1 pb-3 mt-5 mx-3"
                >10 nano</button>
                <button
                  @mouseover="donationAmount=20"
                  @mouseleave="customNanoAmountModel?donationAmount=customNanoAmountModel:donationAmount=2"
                  @click="initiateDonationFor(20)"
                  class="btn w-32 font-medium text-lg bg-black text-white hover:text-lightPink btn-sm-shadow-lightPink rounded-lg px-3 pt-1 pb-3 mt-5 mx-3"
                >20 nano</button>
              </div>
              <!-- Custom Amount Input Group -->
              <form
                v-if="!isDonationInitiated"
                class="w-full md:w-64 flex flex-col justify-center px-3 my-5"
              >
                <label class="w-full text-xl font-bold" for="customNanoAmount">custom amount</label>
                <input
                  :class="inputError?'border-red text-red':'border-black'"
                  class="w-full text-lg font-medium border-2 px-4 pt-1 pb-2 rounded-lg my-1 transition-colors duration-200 ease-out"
                  type="number"
                  ref="customNanoAmount"
                  id="customNanoAmount"
                  name="customNanoAmount"
                  placeholder="enter amount"
                  v-model="customNanoAmountModel"
                  @input="inputChange()"
                />
                <button
                  @click.prevent="customAmountAction()"
                  class="w-full btn text-lg font-medium border-black hover:text-lightPink hover:border-lightPink border-2 bg-black text-white pt-1 pb-2 px-6 rounded-lg my-1"
                >donate</button>
              </form>
              <!-- QR Code for the Donation -->
              <div v-if="isDonationInitiated" class="flex flex-row justify-center items-center m-4">
                <qrcode-vue class="m-4" :value="qrValue" :size="qrSize" level="Q"></qrcode-vue>
                <div class="flex flex-col m-4">
                  <h5 class="text-lg">scan to donate</h5>
                  <h4 class="text-2xl font-bold break-all">{{donationAmount}} nano</h4>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="w-full max-w-lg md:max-w-full px-4 md:px-24 md:hidden mt-12">
        <img
          class="w-full h-auto"
          :src="require('~/assets/images/illustrations/opensource-and-free-mobile.svg')"
          alt="Open-source and Free Mobile"
        />
      </div>
      <div class="w-full md:px-16 lg:px-24 mt-20 hidden md:block">
        <img
          class="w-full h-auto"
          :src="require('~/assets/images/illustrations/opensource-and-free-desktop.svg')"
          alt="Open-source and Free Desktop"
        />
      </div>
    </div>
  </div>
</template>
<script>
import { mapState } from 'vuex'
import QrcodeVue from "qrcode.vue";
import Big from "big.js";
export default {
  data() {
    return {
      isDropdownOpen: false,
      isDonationInitiated: false,
      qrValueBase:
        "nano:nano_1natrium1o3z5519ifou7xii8crpxpk8y65qmkih8e8bpsjri651oza8imdd?amount=",
      qrValue:
        "nano:nano_1natrium1o3z5519ifou7xii8crpxpk8y65qmkih8e8bpsjri651oza8imdd?amount=2000000000000000000000000000000",
      qrSize: 150,
      donationAmount: 2,
      inputError: false
    };
  },
  components: {
    QrcodeVue
  },
  methods: {
    nanoToRaw(inAmount) {
      let nanoRaw = Big(10).pow(30);
      let nanoAmount = Big(inAmount);
      return nanoRaw.times(nanoAmount).toFixed();
    },
    openDonateDropdown() {
      this.isDropdownOpen = true;
    },
    closeDonateDropdown() {
      this.isDropdownOpen = false;
    },
    initiateDonationFor(nanoAmount) {
      this.isDonationInitiated = true;
      this.donationAmount = nanoAmount;
      this.qrValue = this.qrValueBase + this.nanoToRaw(this.donationAmount);
    },
    resetDonation() {
      this.isDonationInitiated = false;
      this.qrValue =
        "nano:nano_1natrium1o3z5519ifou7xii8crpxpk8y65qmkih8e8bpsjri651oza8imdd?amount=2000000000000000000000000000000";
      this.donationAmount = 2;
    },
    customAmountAction() {
      if (this.donationAmount >= 0.000001 && this.donationAmount <= 10000000) {
        this.initiateDonationFor(this.donationAmount);
      } else {
        this.inputError = true;
      }
    },
    inputChange() {
      if (this.inputError) {
        this.inputError = false;
      }
      this.donationAmount = this.$refs.customNanoAmount.value;
    },
    handleAmountCallback(rawAmount) {
      console.log(rawAmount)
    }    
  },
  computed: mapState(['clientID']),
  mounted() {
    this.socket = this.$nuxtSocket({
      name: 'natricon'
    })
    let inst = this
    this.socket.on('connected', function(data) {
      inst.$store.commit("SET_ID", data);
    });
    this.socket.on('donation_event', function(data) {
      inst.handleAmountCallback(data.amount)
    })
  }  
};
</script>
<style scoped>
.btn {
  transition: all 0.2s ease-out;
  transform: scale(1);
}
.btn:hover {
  transform: scale(0.95);
}
.btn-shadow-cyan {
  box-shadow: -0.3rem 0.4rem 0rem 0rem#66FFFF;
}
.btn-sm-shadow-lightPink {
  box-shadow: -0.25rem 0.3rem 0rem 0rem#F199FF;
}
.btn-shadow-lightPink {
  box-shadow: -0.3rem 0.4rem 0rem 0rem#F199FF;
}
.shadow-lightPink {
  box-shadow: -0.5rem 0.6rem 0rem 0rem#F199FF;
}
.btn-shadow-black {
  box-shadow: -0.3rem 0.4rem 0rem 0rem#000000;
}
.btn-shadow-cyan:hover {
  box-shadow: 0rem 0rem 0rem 0rem#66FFFF;
}
.btn-shadow-lightPink:hover,
.btn-sm-shadow-lightPink:hover {
  box-shadow: 0rem 0rem 0rem 0rem#F199FF;
}
.btn-shadow-black:hover {
  box-shadow: 0rem 0rem 0rem 0rem#000000;
}
.line-cyan::after {
  display: block;
  position: absolute;
  width: calc(100% + 0.3rem);
  left: -0.15rem;
  content: "";
  height: 0.75rem;
  border-radius: 0.15rem;
  margin-left: auto;
  margin-right: auto;
  margin-top: -0.9rem;
  background-color: #66ffff;
  z-index: -1;
}
.line-lightPink::after {
  display: block;
  position: absolute;
  width: calc(100% + 0.3rem);
  left: -0.15rem;
  content: "";
  height: 0.75rem;
  border-radius: 0.15rem;
  margin-left: auto;
  margin-right: auto;
  margin-top: -0.9rem;
  background-color: #f199ff;
  z-index: -1;
}
.border-transparent {
  border-color: rgba(0, 0, 0, 0);
}
</style>