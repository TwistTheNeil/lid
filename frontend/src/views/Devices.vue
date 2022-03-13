<template>
  <div class="home">
    <SearchbarInput />
    <DeviceList />
  </div>
</template>

<script>
import SearchbarInput from "@/components/SearchbarInput.vue";
import DeviceList from "@/components/DeviceList.vue";

export default {
  name: "Devices",
  components: {
    SearchbarInput,
    DeviceList,
  },
};
</script>
