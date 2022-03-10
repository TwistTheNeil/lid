import { defineStore } from "pinia";

export const useStorageDeviceStore = defineStore("storageDevices", {
  state: () => ({ devices: [] }),
  actions: {
    async refresh() {
      const response = await fetch(`/api/v1/devices`);
      const responsejson = await response.json();
      this.devices = responsejson;
    },
  },
});
