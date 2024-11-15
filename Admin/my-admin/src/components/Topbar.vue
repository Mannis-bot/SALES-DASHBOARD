<template>
    <div class="top-bar">
      <div class="search-bar">
        <input type="text" v-model="searchQuery" placeholder="Search products...">
      </div>
  
     
      <div class="cart">
        <router-link to="/cart">
          <i class="cart-icon">🛒</i>
        </router-link>
      </div>
  
      
      <div class="profile">
        <div class="profile-info">
          <p class="user-name">Lee Nul</p> 
          <p class="role">Admin</p> 
        </div>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    
    }
  
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
  