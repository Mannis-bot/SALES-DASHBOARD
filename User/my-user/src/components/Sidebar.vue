<template>
    <div class="sidebar">
      <!-- Top links -->
      <div class="top-links">
        <ul>
          <li><router-link to="/" active-class="active-link">Home</router-link></li>
          <li><router-link to="/cart" active-class="active-link">Cart</router-link></li>
          
        </ul>
      </div>
  
      <!-- Bottom logout button -->
      <div class="bottom-logout">
        <button @click="logout">Logout</button>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    methods: {
      logout() {
        
        console.log('User logged out');
      }
    }
  }
  </script>
  
  <style scoped>
  .sidebar {
    display: flex;
    flex-direction: column;
    height: 100vh;  
    width: 250px;   
      
    background-color: white;
    color: black;
    padding: 20px;
    padding-top: 80px;
  }
  
  .top-links {
    flex: 1;  
  }
  
  .top-links ul {
    list-style-type: none;
    padding: 0;
    
  }
  
  .top-links li {
    margin: 20px 0;
    
  }
  
  .top-links a {
    color: black;
    text-decoration: none;
    font-size: 15px;
  }
  
  .bottom-logout {
    margin-top: auto;  
    margin-bottom: 20px;  
  }
  
  .bottom-logout button {
    width: 100%;
    padding: 10px;
    background-color: #e74c3c;
    color: white;
    border: none;
    font-size: 16px;
    cursor: pointer;
  }
  
  .bottom-logout button:hover {
    background-color: #c0392b;
  }
  .active-link {
  color: #72ba14 !important;
}
  </style>
  