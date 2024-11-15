<template>
    <div>
      <form action="" class="search-bar">
        <input type="search" name="" id="" placeholder="Search">
      </form>
      <div class="productsCont mt-[25px] grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-[20px]">
        <div v-for="product in products" :key="product.id" class="product bg-[#ffe8f8] p-[20px] rounded-[15px] border-4 border-white flex flex-col gap-[10px]">
          <img src="../assets/unga.png" alt="" width="100px">
          <h3 class="font-semibold">{{ product.name }}</h3>
          <p class="text-xs font-light">{{ product.short_description }}</p>
          <div class="flex justify-between mt-[15px]">
            <h3 class="font-semibold">Ksh.{{ product.price }}</h3>
            <button @click="addToCart(product)" class="text-[#83ba36] bg-white px-[15px] py-[2px] rounded-md">+ Add</button>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  import { useCartStore } from '@/stores/cartStore'; 
  
  export default {
    data() {
      return {
        products: [], 
      };
    },
    mounted() {
      this.fetchProducts();
    },
    methods: {
      async fetchProducts() {
        try {
          const response = await fetch('http://localhost:8080/api/products');
          const data = await response.json();
          this.products = data;
        } catch (error) {
          console.error('Error fetching products:', error);
        }
      },
      addToCart(product) {
        const cartStore = useCartStore();
        cartStore.addToCart(product); 
        console.log('Product added to cart:', product);
      }
    }
  };
  </script>
  
  
