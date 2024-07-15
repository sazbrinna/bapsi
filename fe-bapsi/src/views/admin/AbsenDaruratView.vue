<template>
  <h1 class="title gap">Absen Darurat</h1>
  <p>Absen ini dilakukan jika terjadi masalah pada bagian absen asisten</p>

  <section>
    <form @submit.prevent="submitData">
      <span>
        <label for="nama">Nama</label>
        <select id="nama" v-model="id_staff">
          <option value="">Pilih asisten...</option>
          <option :value="staff.id_user" v-for="(staff, index) in staffs" :key="index">{{ staff.nama }}</option>
        </select>
      </span>

      <span>
        <label for="region">Loket absen</label>
        <select id="region" v-model="region">
          <option value="">Pilih region...</option>
          <option v-for="(region, index) in regions" :key="index" :value="region">{{ region.region }}</option>
        </select>
      </span>

      <span>
        <label for="jenisAbsen">Jenis Absen</label>
        <select id="jenisAbsen" v-model="jenisAbsen">
          <option value="masuk">Masuk</option>
          <option value="keluar">Keluar</option>
        </select>
      </span>

      <input type="submit" class="bt-1">
    </form>
  </section>
</template>

<script>
import { axios } from "@/axios/config.js";

export default {
  data() {
    return {
      staffs: [],
      regions: [],

      id_staff: "",
      region: [],
      jenisAbsen: "masuk",
    }
  },

  created() {
    this.getStaff();
    this.getRegions();
  },

  methods: {
    getStaff() {
      axios.get("/staffs")
        .then((response) => {
          this.staffs = response.data;
        })
    },

    getRegions() {
      axios.get("/location/all")
        .then((response) => {
          this.regions = response.data;
        })
    },

    submitData() {
      const data = {
        latitude: this.region.latitude.toString(),
        longitude: this.region.longitude.toString(),
        lokasi: this.region.region,
        id_user: this.id_staff,
      }


      if (this.jenisAbsen == "masuk") {
        axios.post("/absen/masuk", data)
        .then(() => {
          window.location.reload();
        })
        .catch((err) => {
          const errorCode = err.response.status
          if (errorCode == 406) {
            this.status = false;
          }
        })
      } else if (this.jenisAbsen == "keluar") {
        axios.post("/absen/keluar", data)
        .then(() => {
          window.location.reload();
        })
        .catch((err) => {
          const errorCode = err.response.status
          if (errorCode == 406) {
            this.status = false;
          }
        })
      }
    },
  }
}
</script>