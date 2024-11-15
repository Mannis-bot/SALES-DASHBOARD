<template>
    <div class="sidebar">
      <!-- Top links -->
      <div class="top-links">
        <ul>
          <li><router-link to="/" active-class="active">Home</router-link></li>
          <li><router-link to="/orders" active-class="active">Orders Management</router-link></li>
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
        // Logic for logging out (clear session, redirect, etc.)
        console.log('User logged out');
      }
    }
  }
  </script>
  
  <style scoped>
  .sidebar {
    display: flex;
    flex-direction: column;
    height: 100vh;  /* Take full height of the viewport */
    width: 250px;   /* Set the width of the sidebar */
      /* Example sidebar background color */
    background-color: white;
    color: black;
    padding: 20px;
    padding-top: 80px;
  }
  
  .top-links {
    flex: 1;  /* This takes up the majority of the space */
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
    margin-top: auto;  /* This pushes the logout button to the bottom */
    margin-bottom: 20px;  /* Optional space from the bottom */
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
  .active {
  color: #72ba14 !important;
}
  </style>
  