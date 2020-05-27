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
          :class="[isDropdownOpen ?'scale-y-100 border-black':'scale-y-0 border-transparent', donationSuccess && isDonationInitiated?'bg-lightGreen':'bg-white', isDropdownOpen && donationSuccess ? 'shadow-black':'', isDropdownOpen && !donationSuccess ?'shadow-lightPink':'']"
          class="w-full md:w-144 absolute flex flex-col items-center max-w-128 border-4 transition-all transform origin-top duration-300 ease-out overflow-hidden mt-5 rounded-lg px-2 md:px-12 z-50"
        >
          <!-- Go Back Button -->
          <button
            v-if="isDonationInitiated"
            class="absolute left-0 top-0 mx-3 my-2 font-bold z-50"
            @click="resetDonation()"
          >
            <div class="relative">
              <div class="w-8 h-8 line bga-lightPink scaleY">
                <img
                  class="w-full h-full"
                  :src="require('~/assets/images/icons/back.svg')"
                  alt="Back Icon"
                />
              </div>
            </div>
          </button>
          <div
            :class="isDropdownOpen?'opacity-100 duration-700':'opacity-0 duration-150' "
            class="w-full flex flex-col justify-center items-center ease-out py-4"
          >
            <img
              v-if="donationSuccess"
              class="w-32 h-32 my-4"
              :src="require('~/assets/images/gifs/NatriconDonatePhase5.gif')"
              alt="Natricon Donate 5"
            />
            <img
              v-else-if="donationAmount>2 && donationAmount<=10"
              class="w-32 h-32 my-4"
              :src="require('~/assets/images/gifs/NatriconDonatePhase2.gif')"
              alt="Natricon Donate 2"
            />
            <img
              v-else-if="donationAmount>10 && donationAmount<100"
              class="w-32 h-32 my-4"
              :src="require('~/assets/images/gifs/NatriconDonatePhase3.gif')"
              alt="Natricon Donate 3"
            />
            <img
              v-else-if="donationAmount>=100"
              class="w-32 h-32 my-4"
              :src="require('~/assets/images/gifs/NatriconDonatePhase4.gif')"
              alt="Natricon Donate 4"
            />
            <img
              v-else
              class="w-32 h-32 my-4"
              :src="require('~/assets/images/gifs/NatriconDonatePhase1.gif')"
              alt="Natricon Donate 1"
            />
            <div class="flex flex-row justify-center">
              <!-- Donation Initated -->
              <div v-if="!isDonationInitiated" class="flex flex-col items-center">
                <!-- Donate Amount Buttons -->
                <div class="flex flex-row flex-wrap justify-center items-center my-4">
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
                <form class="w-full md:w-64 flex flex-col justify-center px-3 my-5">
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
              </div>
              <!-- QR Code for the Donation -->
              <div
                v-else-if="isDonationInitiated  && !donationSuccess"
                class="flex flex-row flex-wrap justify-center items-center m-4"
              >
                <div class="mx-4 my-4 border-4 rounded-lg p-1 border-lightPink bg-white qr-shadow">
                  <qrcode-vue :value="qrValue" :size="qrSize" level="Q"></qrcode-vue>
                </div>
                <div class="flex flex-col m-4">
                  <h5 class="text-lg px-2">scan to donate</h5>
                  <h4 class="text-2xl font-bold break-all px-2">{{donationAmount}} nano</h4>
                  <button
                    @click="doCopy()"
                    ref="copyButton"
                    :class="isAddressCopied?'bg-lightPink':'bg-white'"
                    class="text-xs font-mono hover:bg-lightPink text-left mt-3 rounded-lg transition-colors duration-300 ease-out p-2"
                    v-html="isAddressCopied?copiedHtml:addressHtml"
                  ></button>
                </div>
              </div>
              <!-- Success Screen Text -->
              <div
                class="flex flex-col items-center my-4"
                v-else-if="isDonationInitiated && donationSuccess"
              >
                <h4 class="text-4xl text-center font-bold">thank you!</h4>
                <h5 v-if="donationAmount>=2" class="text-lg text-center">your badge is on its way...</h5>
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
import { mapState } from "vuex";
import QrcodeVue from "qrcode.vue";
import Big from "big.js";
import Vue from "vue";
import VueClipboard from "vue-clipboard2";

const donationAddress =
  "nano_1natrium1o3z5519ifou7xii8crpxpk8y65qmkih8e8bpsjri651oza8imdd";

VueClipboard.config.autoSetContainer = true; // add this line
Vue.use(VueClipboard);
export default {
  data() {
    return {
      address: donationAddress,
      isDropdownOpen: false,
      isDonationInitiated: false,
      donationAmountModifierBase: 0.001,
      qrValueAmountRaw: "",
      qrValue: "",
      qrSize: 150,
      donationAmount: 2,
      inputError: false,
      donationSuccess: false,
      customNanoAmountModel: null,
      isAddressCopied: false,
      addressHtml: `${donationAddress.substring(
        0,
        22
      )}<br/>${donationAddress.substring(
        22,
        44
      )}<br/>${donationAddress.substring(44, 65)}`,
      copiedHtml: `&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp<br/>&nbsp&nbsp&nbsp&nbspaddress copied&nbsp&nbsp&nbsp&nbsp<br/>&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp`
    };
  },
  components: {
    QrcodeVue
  },
  methods: {
    nanoToRaw(inAmount) {
      let nanoRaw = Big(10).pow(30);
      let nanoAmount = Big(inAmount);
      return nanoRaw.times(nanoAmount);
    },
    appendIdToRaw(inAmount) {
      // Modifify donation amount with 0.001 + socketio client ID
      // Will let us recognize this donation
      let idModifier = Big(this.clientID);
      let amountModifier = Big(10).pow(27); // 0.001 NANO
      return inAmount
        .add(idModifier)
        .add(amountModifier)
        .toFixed();
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
      this.qrValueAmountRaw = this.appendIdToRaw(
        this.nanoToRaw(this.donationAmount)
      );
      this.qrValue = `nano:${donationAddress}?amount=${this.qrValueAmountRaw}`;
    },
    resetDonation() {
      this.isDonationInitiated = false;
      this.qrValue = "";
      this.qrValueAmountRaw = "";
      this.donationAmount = 2;
      this.donationSuccess = false;
    },
    customAmountAction() {
      this.donationAmount = this.$refs.customNanoAmount.value;
      if (this.donationAmount >= 0.000001 && this.donationAmount <= 10000000) {
        this.initiateDonationFor(Number(this.donationAmount));
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
    doCopy() {
      let ref = this;
      this.$copyText(this.address).then(
        function(e) {
          ref.isAddressCopied = true;
        },
        function(e) {
          alert("Can not copy");
        }
      );
      setTimeout(function() {
        ref.isAddressCopied = false;
      }, 2000);
    },
    handleAmountCallback(rawAmount) {
      if (rawAmount == this.qrValueAmountRaw) {
        this.donationSuccess = true;
      }
    }
  },
  computed: mapState(["clientID"]),
  mounted() {
    this.socket = this.$nuxtSocket({
      name: "natricon"
    });
    let inst = this;
    this.socket.on("connected", function(data) {
      // Use ID sent from server as a donation modifier
      inst.$store.commit("SET_ID", data);
    });
    this.socket.on("donation_event", function(data) {
      inst.handleAmountCallback(data.amount);
    });
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
.shadow-black {
  box-shadow: -0.5rem 0.6rem 0rem 0rem#000000;
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
.qr-shadow {
  box-shadow: -0.3rem 0.3rem 0rem 0rem#000000;
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
.bga-lightPink::after {
  background-color: #f199ff;
}
.line::after {
  display: block;
  position: absolute;
  width: calc(100% + 0.4rem);
  left: -0.2rem;
  content: "";
  height: 1rem;
  border-radius: 0.15rem;
  margin-left: auto;
  margin-right: auto;
  margin-top: -1rem;
  z-index: -1;
  transition: all 0.2s ease-out;
  transform-origin: center bottom;
  transform: scaleY(0);
}
.scaleY:hover::after {
  transform: scaleY(1);
}
</style>