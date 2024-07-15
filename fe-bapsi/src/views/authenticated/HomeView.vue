<template>
  <HeroImage />

  <section class="top bottom">
    <div class="wrapper">
      <h1 class="title">Kehadiran di Ruangan</h1>

      <ProfileHadir :users="id_user" v-if="id_user" />
      <p v-else>Tidak ada yang sedang berada di loket</p>

    </div>

  </section>

  <AlertCard :msg="message" v-if="message" />
</template>

<script>
import { axios } from "@/axios/config.js";
import HeroImage from "@/components/hero-image/HeroImage.vue";
import ProfileHadir from "@/components/small-part/wrapper/WrapperCardProfile.vue";
import AlertCard from "@/components/alert/AlertCard.vue";

export default {
  components: {
    HeroImage,
    ProfileHadir,
    AlertCard,
  },

  data() {
    return {
      id_user: [],

      message: "",
    }
  },

  beforeCreate() {
    axios.get("/absen/cek/masuk")
      .then((response) => {
        var emp = response.data;
        this.id_user = emp;
      })
      .catch(() => {
        this.message = "Gagal menerima input";
            setTimeout(() => {
              this.message = "";
            }, 2000)
      })
  }
}
</script>

<style scoped>
span {
  display: flex;
  width: 100%;
}
</style>