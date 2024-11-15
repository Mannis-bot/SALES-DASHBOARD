<template>
    <div>
        <form action="" class="search-bar">
            <input type="search" name="" id="" placeholder="Search">
        </form>
        <div class="productsCont mt-[25px]">
            <div class="product w-[230px] bg-[#ffe8f8] p-[20px] rounded-[15px] border-4 border-white flex flex-col gap-[10px]">
            <img src="../assets/unga.png" alt="" width="100px">
            <h3 class="font-semibold">Ajab Baking Flour</h3>
            <p class="text-xs font-light">Flour 2Kg , also known as plain flour is commonly used for 
                baking purposes. The protein level lies between 10 to 12 percent.</p>
            <div class="flex justify-between mt-[15px]">
                <h3 class="font-semibold">Ksh.198</h3>
                <button class="text-[#83ba36] bg-white px-[15px] py-[2px] rounded-md">+ Add</button>
            </div>
        </div>
        </div>
    </div>
</template>

<script>
    export default {
        
    }
</script>

<style  scoped>
 .search-bar input {
    padding: 8px 15px;
    border-radius: 25px;
    width: 250px;
  }

</style>