<template>
  <div class="order-management-container">
    <h2 class="text-xl font-bold mb-4">Order Management</h2>

    <div v-if="loading" class="spinner">
      <p>Loading orders...</p>
    </div>

    <div v-else>
      <table class="order-table w-full table-auto border-collapse">
        <thead>
          <tr class="bg-gray-100">
            <th class="py-2 px-4 border">User Name</th>
            <th class="py-2 px-4 border">Total Price</th>
            <th class="py-2 px-4 border">Status</th>
            <th class="py-2 px-4 border">Actions</th>
          </tr>
        </thead>
        <tbody>
          <!-- Loop through grouped orders by user_name -->
          <tr v-for="(userOrders, userName) in groupedOrders" :key="userName">
            <td class="py-2 px-4 border text-center" rowspan="userOrders.orders.length">{{ userName }}</td>
            <td class="py-2 px-4 border text-center">
              Ksh {{ userOrders.totalPrice.toFixed(2) }}
            </td>
            <td class="py-2 px-4 border text-center">
              <span :class="{
                'text-green-600': userOrders.status === 'completed',
                'text-yellow-500': userOrders.status === 'pending'
              }">{{ userOrders.status }}</span>
            </td>
            <td class="py-2 px-4 border text-center">
              <button @click="showOrderDetails(userOrders.orders)" class="bg-blue-500 text-white py-1 px-4 rounded">View More</button>
            </td>
          </tr>
        </tbody>
      </table>

      <!-- Order Details Modal -->
      <div v-if="orderDetails" class="fixed inset-0 bg-gray-800 bg-opacity-50 flex justify-center items-center z-50">
        <div class="bg-white rounded-lg shadow-lg w-96 p-6">
          <h3 class="text-2xl font-semibold text-center mb-6">Order Details</h3>
          
          <ul class="space-y-4 text-lg">
            <li>
              <span class="font-medium">User Name:</span> {{ orderDetails.userName }}
            </li>
            <li v-for="(order, index) in orderDetails.orders" :key="index">
              <span class="font-medium">Product Name:</span> {{ order.product_name }}<br />
              <span class="font-medium">Price:</span> Ksh {{ order.product_price.toFixed(2) }}
            </li>
            <li>
              <span class="font-medium">Total Price:</span> Ksh {{ orderDetails.totalPrice.toFixed(2) }}
            </li>
            <li v-if="orderDetails.comment">
              <span class="font-medium">Comment:</span> {{ orderDetails.comment }}
            </li>
          </ul>

          <div class="flex justify-end mt-6">
            <button @click="closeOrderDetails" class="bg-blue-500 text-white py-2 px-6 rounded-lg hover:bg-blue-700 transition duration-200">
              Close
            </button>
          </div>
        </div>
      </div>

      <div v-if="orders.length > 10" class="mt-4 flex justify-between">
        <button class="btn-pagination" @click="previousPage">Previous</button>
        <button class="btn-pagination" @click="nextPage">Next</button>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, onMounted, computed } from "vue";
import axios from "axios";

export default {
  name: "OrderManagement",
  setup() {
    const orders = ref([]);
    const loading = ref(true);
    const orderDetails = ref(null);
    const page = ref(1);
    const limit = ref(10);

    // Fetch orders from the API
    const fetchOrders = async () => {
      try {
        loading.value = true;
        const url = `http://localhost:8080/api/orders?page=${page.value}&limit=${limit.value}`;
        const response = await axios.get(url);
        orders.value = response.data;
      } catch (error) {
        console.error("Error fetching orders:", error);
      } finally {
        loading.value = false;
      }
    };

    // Show order details in the modal
    const showOrderDetails = (userOrders) => {
      const totalPrice = userOrders.reduce((sum, order) => sum + order.product_price, 0);
      orderDetails.value = {
        userName: userOrders[0].user_name,
        orders: userOrders,
        totalPrice: totalPrice,
      };
    };

    // Close the order details modal
    const closeOrderDetails = () => {
      orderDetails.value = null;
    };

    // Pagination: next page
    const nextPage = () => {
      page.value += 1;
      fetchOrders();
    };

    // Pagination: previous page
    const previousPage = () => {
      if (page.value > 1) {
        page.value -= 1;
        fetchOrders();
      }
    };

    // Group orders by user name and calculate total price per user
    const groupedOrders = computed(() => {
      const groups = {};
      orders.value.forEach((order) => {
        if (!groups[order.user_name]) {
          groups[order.user_name] = { orders: [], totalPrice: 0, status: order.status };
        }
        groups[order.user_name].orders.push(order);
        groups[order.user_name].totalPrice += order.product_price;
      });
      return groups;
    });

    // Initialize orders
    onMounted(() => {
      fetchOrders();
    });

    return {
      orders,
      loading,
      orderDetails,
      showOrderDetails,
      closeOrderDetails,
      nextPage,
      previousPage,
      groupedOrders,
    };
  },
};
</script>


<style scoped>
.order-management-container {
  margin: 20px;
}

.order-table {
  border-collapse: collapse;
}

.order-table th,
.order-table td {
  text-align: left;
  padding: 10px;
}

.order-table td {
  border-bottom: 1px solid #ddd;
}

.spinner {
  text-align: center;
  font-size: 18px;
  color: #666;
}

.btn-pagination {
  background-color: #007bff;
  color: white;
  padding: 8px 16px;
  border-radius: 5px;
  cursor: pointer;
}

.btn-pagination:hover {
  background-color: #0056b3;
}

.order-details-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
}

.order-details-modal {
  background: white;
  padding: 20px;
  border-radius: 8px;
  width: 400px;
  max-height: 80vh;
  overflow-y: auto;
}
</style>


