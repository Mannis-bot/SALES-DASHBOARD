<template>
    <div class="dashboard-container">
      
      <div class="welcome-section p-6 bg-white shadow-lg rounded-lg mb-6">
        <h2 class="text-2xl font-semibold text-green-600">Welcome to the Admin Dashboard</h2>
        <p class="text-gray-600 mt-2">Manage your store, view orders, and update products.</p>
      </div>
  
      
      <div class="stats grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
        <div class="stat-card bg-white p-6 shadow-lg rounded-lg">
          <h3 class="text-xl font-semibold text-gray-800">Total Sales</h3>
          <p class="text-2xl font-bold text-green-600">Ksh 1,200,000</p>
        </div>
        <div class="stat-card bg-white p-6 shadow-lg rounded-lg">
          <h3 class="text-xl font-semibold text-gray-800">Products</h3>
          <p class="text-2xl font-bold text-green-600">150</p>
        </div>
        <div class="stat-card bg-white p-6 shadow-lg rounded-lg">
          <h3 class="text-xl font-semibold text-gray-800">Orders</h3>
          <p class="text-2xl font-bold text-green-600">350</p>
        </div>
      </div>
  
      
      <div class="chart mt-6 p-6 bg-white shadow-lg rounded-lg">
        <h3 class="text-xl font-semibold text-gray-800">Sales Overview</h3>
       
        <LineChart :data="chartData" :options="chartOptions" />
      </div>
    </div>
  </template>
  
  <script>
  import { defineComponent } from "vue";
  import { Line } from "vue-chartjs";
  import {
    Chart as ChartJS,
    Title,
    Tooltip,
    Legend,
    LineElement,
    PointElement,
    CategoryScale,
    LinearScale,
  } from "chart.js";
  
  
  ChartJS.register(Title, Tooltip, Legend, LineElement, PointElement, CategoryScale, LinearScale);
  
  export default defineComponent({
    name: "Dashboard",
    components: {
      LineChart: Line,
    },
    data() {
      return {
        chartData: {
          labels: ["January", "February", "March", "April", "May"], 
          datasets: [
            {
              label: "Sales",
              data: [100, 200, 150, 300, 250], 
              fill: true,
              backgroundColor: "rgba(72, 187, 120, 0.2)",
              borderColor: "rgba(72, 187, 120, 1)",
              borderWidth: 2,
            },
          ],
        },
        chartOptions: {
          responsive: true,
          plugins: {
            title: {
              display: true,
              text: "Monthly Sales Overview",
            },
          },
          scales: {
            x: {
              beginAtZero: true,
            },
            y: {
              beginAtZero: true,
            },
          },
        },
      };
    },
  });
  </script>
  
  <style scoped>
  .dashboard-container {
    padding: 20px;
  }
  
  .welcome-section {
    background-color: #ffffff;
  }
  
  .stats {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    gap: 1.5rem;
  }
  
  .stat-card {
    background-color: #ffffff;
    padding: 1.5rem;
    border-radius: 10px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  }
  
  .stat-card h3 {
    color: #2d3748;
  }
  
  .stat-card p {
    color: #48bb78;
    font-size: 1.5rem;
    font-weight: 700;
  }
  
  .chart {
    background-color: #ffffff;
    padding: 1.5rem;
    border-radius: 10px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  }
  
  .sales-graph {
    padding-top: 1rem;
    background-color: #f7fafc;
  }
  </style>
  
  