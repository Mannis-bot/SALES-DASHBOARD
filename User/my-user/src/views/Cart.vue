<template>
  <div class="min-h-screen bg-gray-50">
    <div class="container mx-auto p-6">
      <!-- Cart Items -->
      <div v-if="cartItems.length > 0" class="bg-white p-6 shadow-lg rounded-lg mb-6">
        <h2 class="text-lg font-semibold text-green-600">Your Cart</h2>
        <ul class="space-y-4 mt-4">
          <li v-for="(item, index) in cartItems" :key="index" class="flex justify-between">
            <div>
              <span class="text-gray-800">{{ item.name }}</span>
              <div class="mt-2 flex items-center space-x-2">
                <button @click="updateQuantity(item, -1)" class="btn">-</button>
                <span>{{ item.quantity }}</span>
                <button @click="updateQuantity(item, 1)" class="btn">+</button>
              </div>
            </div>
            <span class="text-gray-600">{{ (item.price * item.quantity) | currency }} Ksh</span>
          </li>
        </ul>
        
        <!-- Subtotal and Tax -->
        <div class="mt-4">
          <div class="flex justify-between">
            <span class="text-gray-700">Subtotal:</span>
            <span class="text-gray-600">{{ subtotal | currency }} Ksh</span>
          </div>
          <div class="flex justify-between">
            <span class="text-gray-700">Tax (10%):</span>
            <span class="text-gray-600">{{ tax | currency }} Ksh</span>
          </div>
          <div class="flex justify-between font-semibold">
            <span>Total:</span>
            <span class="text-gray-600">{{ total | currency }} Ksh</span>
          </div>
        </div>

        <!-- Comment Section -->
        <div class="mt-4">
          <textarea v-model="orderComment" placeholder="Add a comment (e.g., deliver by 24 hours)" class="w-full p-2 border border-gray-300 rounded"></textarea>
        </div>
      </div>
      <div v-else class="text-center text-gray-600">Your cart is empty.</div>

      <!-- Checkout Button -->
      <div v-if="cartItems.length > 0" class="mt-4">
        <button @click="checkout" class="btn-green">Place Order</button>
      </div>
    </div>
  </div>
</template>

<script>
import { useCartStore } from "@/stores/cartStore";
import { useRouter } from "vue-router";

export default {
  data() {
    return {
      orderComment: '',
    };
  },
  computed: {
    cartItems() {
      const cartStore = useCartStore();
      return cartStore.cartItems;
    },
    subtotal() {
      return this.cartItems.reduce((total, item) => total + item.price * item.quantity, 0);
    },
    total() {
      return this.subtotal 
    }
  },
  methods: {
    updateQuantity(item, change) {
      const cartStore = useCartStore();
      const newQuantity = item.quantity + change;
      if (newQuantity > 0) {
        cartStore.updateQuantity(item.id, newQuantity);
      } else {
        cartStore.removeFromCart(item.id);
      }
    },
    async checkout() {
      const username = "postgres"; 
      const password = "Marykiki";   //My Password

      
      const encodedCredentials = btoa(`${username}:${password}`);

      try {
        const orderData = {
          items: this.cartItems.map(item => ({
            product_id: item.id,
            quantity: item.quantity,
          })),
          user_name: "DemoUser", 
          total_price: this.total,
          tax: this.tax,
          comment: this.orderComment
        };
        console.log("Order Data:", orderData);

        const response = await fetch("http://localhost:8080/api/place-order", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            "Authorization": `Basic ${encodedCredentials}` 
          },
          body: JSON.stringify(orderData),
        });

        if (response.ok) {
          alert("Order placed successfully!");
          const cartStore = useCartStore();
          console.log()
          cartStore.clearCart(); // Clear cart after successful checkout
          this.$router.push("/"); // Redirect to home page
        } else {
          alert("Failed to place order.");
        }
      } catch (error) {
        console.error("Checkout error:", error);
      }
    },
  },
  filters: {
    currency(value) {
      return value ? value.toFixed(2) : "0.00";
    },
  },
};
</script>



