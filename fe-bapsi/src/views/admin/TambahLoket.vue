<template>
  <h1 class="title gap">Tambah Loket</h1>

  <section>
    <!-- sementara edit lokasi langsung di sini -->

    <form @submit.prevent="SubmitData">
      <span>
        <label for="loket">Loket</label>
        <input type="text" id="loket" placeholder="Region..." v-model="region" required />
      </span>

      <span>
        <label for="latitude">Latitude</label>
        <input type="text" id="latitude" placeholder="Latitude..." v-model="latitude" required />
      </span>

      <span>
        <label for="longitude">Longitude</label>
        <input type="text" id="longitude" placeholder="Longitude..." v-model="longitude" required />
      </span>

      <input type="submit" class="bt-1">
    </form>
  </section>

  <AlertCard v-if="message" :msg="message" />
</template>

<script>
import AlertCard from "@/components/alert/AlertCard.vue";
import {axios} from "@/axios/config.js";

export default {
  data() {
    return {
      region: "",
      latitude: "",
      longitude: "",

      message: "",
    }
  },

  components: {
    AlertCard,
  },

  methods: {
    SubmitData() {
      const data = {
        loket: this.region,
        latitude: this.latitude,
        longitude: this.longitude,
      }

      axios.post("/region", data)
      .then(() => {
        this.$router.push({name: 'KonfigurasiView'});
      })
      .catch(() => {
          this.message = "Gagal menambahkan region";
          setTimeout(() => {
            this.message = "";
          }, 2000)
        });
    }
  }
}
</script>