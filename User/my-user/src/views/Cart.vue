<template>
    <div class="min-h-screen bg-gray-50">
      <!-- Top Bar -->
      <div class="bg-green-600 text-white p-4 shadow-md">
        <h1 class="text-xl font-semibold">Your Cart</h1>
      </div>
  
      <!-- Cart Content -->
      <div class="container mx-auto p-6">
        <!-- Cart Items -->
        <div class="bg-white p-4 shadow-lg rounded-lg mb-6">
          <h2 class="text-lg font-semibold text-green-600">Items in your cart</h2>
          <div v-if="cartItems.length > 0" class="mt-4">
            <ul class="space-y-4">
              <li v-for="(item, index) in cartItems" :key="index" class="flex items-center justify-between border-b py-2">
                <span class="text-gray-800">{{ item.name }}</span>
                <span class="text-gray-600">{{ item.price | currency }} Ksh</span>
              </li>
            </ul>
          </div>
          <div v-else class="text-center text-gray-600 mt-6">
            Your cart is empty.
          </div>
        </div>
  
        <!-- Cart Summary -->
        <div class="bg-white p-4 shadow-lg rounded-lg">
          <h2 class="text-lg font-semibold text-green-600">Cart Summary</h2>
          <div class="mt-4 flex justify-between">
            <span class="text-gray-700">Subtotal</span>
            <span class="text-gray-700">{{ subtotal | currency }} Ksh</span>
          </div>
          <div class="mt-2 flex justify-between">
            <span class="text-gray-700">Discount</span>
            <span class="text-gray-700">{{ discount }} %</span>
          </div>
          <div class="mt-4 flex justify-between text-xl font-semibold text-green-600">
            <span>Total</span>
            <span>{{ total | currency }} Ksh</span>
          </div>
  
          <!-- Continue Shopping Button -->
          <div class="mt-6 flex gap-4 justify-between">
            <router-link to="/">
              <button class="py-2 bg-gray-300 text-black rounded-lg hover:bg-gray-400 focus:outline-none">
                Continue Shopping
              </button>
            </router-link>
  
            <!-- Checkout Button -->
            <button
              v-if="!isCheckoutReady"
              class="py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 focus:outline-none"
              @click="proceedToCheckout"
            >
              Proceed to Checkout in {{ countdown }} seconds
            </button>
  
            <!-- After countdown, the button text changes to "Go to Checkout Now" -->
            <button
              v-else
              class="py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 focus:outline-none"
              @click="proceedToCheckout"
            >
              Go to Checkout Now
            </button>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        cartItems: [
          { name: "Sugar", price: 10 },
          { name: "Maize Flour", price: 5 },
          // Add more items as needed
        ],
        discount: 10, // Example discount
        countdown: 5, // Initial countdown time (in seconds)
        isCheckoutReady: false, // Flag to show "Go to Checkout Now" button
      };
    },
    computed: {
      subtotal() {
        return this.cartItems.reduce((acc, item) => acc + item.price, 0);
      },
      total() {
        return this.subtotal - (this.subtotal * this.discount) / 100;
      },
    },
    methods: {
      proceedToCheckout() {
        // Proceed to checkout logic
        this.$router.push("/checkout"); // Assuming you have a checkout route
      },
    },
    mounted() {
      // Start the countdown timer when the component is mounted
      this.startCountdown();
    },
    methods: {
      startCountdown() {
        const countdownInterval = setInterval(() => {
          if (this.countdown > 1) {
            this.countdown--;
          } else {
            this.isCheckoutReady = true;
            clearInterval(countdownInterval);
          }
        }, 1000); // Decrease countdown every second
      },
      proceedToCheckout() {
        // Handle the checkout logic
        this.$router.push("/checkout");
      },
    },
    filters: {
      currency(value) {
        if (!value) return '0';
        return value.toFixed(2); // Ensure the value always has two decimal places
      },
    },
  };
  </script>
  
  <style scoped>
  /* You can customize this further, but Tailwind should be enough for most use cases */
  </style>
  