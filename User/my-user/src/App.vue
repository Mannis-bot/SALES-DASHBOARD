<!-- src/App.vue -->
<template>
  <div class="app-layout">
    <!-- Sidebar (Left navigation) -->
    <Sidebar />

    <!-- Main content area, includes Topbar and dynamic page content -->
    <div class="main-content bg-[#ebf2f0]">
      <!-- Topbar (Top navigation) -->
      <Topbar />

      <!-- Dynamic page content loaded here -->
      <router-view />
    </div>
  </div>
</template>

<script>
import Sidebar from './components/Sidebar.vue';
import Topbar from './components/Topbar.vue';

export default {
  components: {
    Sidebar,
    Topbar,
  },
};
</script>

<style>

.app-layout {
  display: flex;          
  height: 100vh;          
}

/* Sidebar setup */
.sidebar {
  width: 250px;           /* Adjust based on your design */
  background-color: #333; /* Color for sidebar */
  color: white;           /* Text color */
  position: fixed;        /* Fix sidebar on the left */
  top: 0;
  left: 0;
  bottom: 0;
  padding-top: 20px;
  z-index: 1;
}

/* Main content area */
.main-content {
  display: flex;
  flex-direction: column; /* Stack elements vertically (Topbar + router-view) */
  flex: 1;                /* Take up remaining space */
  margin-left: 250px;     /* Offset for sidebar */
  padding: 20px;          /* Padding for content */
  overflow-y: auto;       /* Allow scrolling if content is tall */
}

/* Topbar layout */
.topbar {
  background-color: #444;
  padding: 15px;
  color: white;
}
</style>

