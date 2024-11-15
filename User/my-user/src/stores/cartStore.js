import { defineStore } from 'pinia';

export const useCartStore = defineStore('cart', {
  state: () => ({
    cartItems: [], // Store products added to the cart
  }),
  actions: {
    // Add product to the cart
    addToCart(product) {
      const existingProduct = this.cartItems.find(
        (item) => item.id === product.id
      );
      if (existingProduct) {
        existingProduct.quantity += 1; // Increase quantity if product exists
      } else {
        this.cartItems.push({ ...product, quantity: 1 }); // Add new product
      }
    },
    // Remove product from cart
    removeFromCart(productId) {
      this.cartItems = this.cartItems.filter(item => item.id !== productId);
    },
    // Update product quantity
    updateQuantity(product, newQuantity) {
      const cartItem = this.cartItems.find(item => item.id === product.id);
      if (cartItem) {
        cartItem.quantity = newQuantity;
      }
    },
    // Get the total price of all items in the cart
    getTotalPrice() {
      return this.cartItems.reduce((total, item) => total + item.price * item.quantity, 0);
    },
    // Clear all items in the cart
    clearCart() {
      this.cartItems = [];  // Reset cart items to empty
    }
  },
  getters: {
    itemCount: (state) => state.cartItems.length,
  }
});
