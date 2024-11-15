<template>
    <div class="min-h-screen bg-gray-50">
      <!-- Top Bar -->
      <div class="bg-green-600 text-white p-4 shadow-md">
        <h1 class="text-xl font-semibold">Checkout</h1>
      </div>
  
      <!-- Checkout Content -->
      <div class="container mx-auto p-6">
        <!-- Order Summary -->
        <div class="bg-white p-6 shadow-lg rounded-lg mb-6">
          <h2 class="text-lg font-semibold text-green-600">Order Summary</h2>
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
  
        <!-- Order Confirmation Button -->
        <div class="text-center">
          <button
            @click="placeOrder"
            class="py-2 px-4 bg-green-600 text-white rounded-lg hover:bg-green-700"
          >
            Place Order
          </button>
        </div>
  
        <!-- Success Message (Initially Hidden) -->
        <div v-if="orderPlaced" class="mt-6 p-4 text-green-800 bg-green-200 rounded-lg text-center">
          Your Order has been placed successfully!
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
        orderPlaced: false,
      };
    },
    computed: {
      subtotal() {
        return this.cartItems.reduce((acc, item) => acc + item.price, 0);
      },
      total() {
        return this.subtotal; // You can add discount logic here if needed
      },
    },
    methods: {
      placeOrder() {
        // Show confirmation message
        this.orderPlaced = true;
  
        // Redirect to Home after a short delay (for user experience)
        setTimeout(() => {
          this.$router.push("/"); // Redirect to Home page after order is placed
        }, 2000); // Redirects after 2 seconds
      },
    },
    filters: {
      currency(value) {
        if (!value) return '0';
        return value.toFixed(2);
      },
    },
  };
  </script>
  
  <style scoped>
  /* Add any specific styles you need for this component */
  </style>
  