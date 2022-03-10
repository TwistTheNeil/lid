import { defineStore } from "pinia";

export const useStorageDeviceStore = defineStore("storageDevices", {
  state: () => ({ devices: [] }),
});
