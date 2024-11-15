<template>
    <div class="top-bar">
      <!-- Page Name (Dynamic) -->
      <div class="page-name">
        <h1>{{ pageTitle }}</h1>
      </div>
  
      <!-- Search Bar -->
      <div class="search-bar">
        <input type="text" v-model="searchQuery" placeholder="Search products...">
      </div>
  
      <!-- Cart -->
      <div class="cart">
        <router-link to="/cart">
          <i class="cart-icon">🛒</i>
        </router-link>
      </div>
  
      <!-- Profile -->
      <div class="profile">
        <div class="profile-info">
          <p class="user-name">John Doe</p> <!-- User's Name -->
          <p class="role">User</p> <!-- User Role -->
        </div>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        pageTitle: '',  // Holds the name of the current page
        searchQuery: '', // For search input
      };
    },
    created() {
      // Set the page title when the component is created
      this.updatePageTitle(this.$route.name);
    },
    watch: {
      // Watch for route changes and update the title dynamically
      '$route'(to) {
        this.updatePageTitle(to.name);
      }
    },
    methods: {
      // Update page title based on route name
      updatePageTitle(routeName) {
        if (routeName === 'home') {
          this.pageTitle = 'Home';
        } else if (routeName === 'cart') {
          this.pageTitle = 'Your Cart';
        } else if (routeName === 'order-history') {
          this.pageTitle = 'Order History';
        } 
      }
    }
  };
  </script>
  
  <style scoped>
  .top-bar {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 10px 20px;
    color: black;
    margin-bottom: 20px;
    
  }
  
  .page-name h1 {
    font-size: 1.5rem;
  }
  
  .search-bar input {
    padding: 8px 15px;
    border-radius: 25px;
    width: 250px;
  }
  
  .cart, .profile {
    display: flex;
    align-items: center;
    margin-left: 20px;
  }
  
  .cart-icon {
    font-size: 1rem;
    cursor: pointer;
    background: white;
    border-radius: 50%;
    padding: 5px;
  }
  
  .profile-info {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    margin-left: 10px;
  }
  
  .user-name {
    font-weight: bold;
    font-size: 1.1rem;
  }
  
  .role {
    font-size: 0.9rem;
    color: #666;
  }
  </style>
  
  