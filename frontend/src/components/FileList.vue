<template>
  <table class="table table-striped table-hover table-sm">
    <thead>
      <tr>
        <th v-for="header in headers" scope="col" :key="header">
          {{ header }}
        </th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="item in items" :key="item.uuid || item.hash">
        <td>{{ item.hash }}</td>
        <td>{{ item.name }}</td>
      </tr>
    </tbody>
  </table>
</template>

<script>
export default {
  name: "FileList",
  props: {
    headers: Array,
    items: Array,
  },
};
</script>
