<template>
  <section class="full">
    <div class="wrapper">
      <h2>Registrasi</h2>
      <span>
        <h3>Bapsi Universitas Gunadarma</h3>
      </span>

      <span :class="{none: !registered}">
        <h2 class="green">Permintaan membuat akun baru telah berhasil dilakukan. Silahkan keluar dari website dan tunggu konfirmasi lebih lanjut menlalui Email.</h2>
      </span>

      <form @submit.prevent="submitData" :class="{none: registered}">
        <span>
          <label for="nama">Nama</label>
          <input type="text" id="nama" placeholder="Nama..." v-model="nama" required>
        </span>

        <span>
          <label for="email">Email</label>
          <input type="email" id="email" placeholder="Email..." v-model="email" required>
        </span>

        <span>
          <label for="region">Lokasi Loket</label>
          <select id="region" v-model="id_region">
            <option value="">Pilih region...</option>
            <option v-for="(region, index) in regions" :key="index" :value="region.id_region">{{ region.region }}</option>
          </select>
        </span>

        <span>
          <label for="divisi">Divisi</label>
          <select id="divisi" v-model="id_divisi">
            <option value="">Pilih divisi...</option>
            <option v-for="(divisi, index) in divisis" :key="index" :value="divisi.id_divisi"> {{ divisi.divisi }}
            </option>
          </select>
        </span>

        <span>
          <label for="password">Password</label>
          <input type="password" id="password" placeholder="Password..." v-model="password" required>
          <p class="warning" v-if="password.length < 6 && password.length > 0">Password harus memiliki setidaknya 6 digit</p>
        </span>

        <span>
          <label for="confirm-password">Konfirmasi password</label>
          <input type="password" id="confirm-password" placeholder="Konfirmasi password..." v-model="confirmPassword" required>
          <p class="warning" v-if="confirmPassword.length > 0 && password !== confirmPassword">Password tidak cocok</p>
        </span>

        <a href="" @click.prevent="PushRoute('LoginView')">Sudah memiliki akun?</a>

        <input type="submit" class="bt-1" value="Registrasi">
      </form>
    </div>
  </section>

  <AlertCard v-if="message" :msg="message" />
</template>

<script>
import AlertCard from "@/components/alert/AlertCard.vue";
import { axios } from "@/axios/config.js";

export default {
  data() {
    return {
      regions: [],
      divisis: [],

      id_region: "",
      id_divisi: "",
      role: "asisten",
      nama: "",
      email: "",
      password: "",
      confirmPassword: "",
      statusPassword: false,
      message: "",

      registered: false,
    }
  },

  components: {
    AlertCard,
  },

  created() {
    this.getRegions();
  },

  beforeMount() {
    if (localStorage.getItem("tokenAuth")) {
      this.$router.push({ "name": "HomeView" })
    }
  },

  methods: {
    PushRoute(pageName) {
      window.scrollTo({
        top: 0,
        behavior: 'smooth'
      });
      
      this.$router.push({ name: pageName })
    },

    getRegions() {
      axios.get("/location/all")
        .then((response) => {
          this.regions = response.data;
        })

      axios.get("/divisi")
        .then((response) => {
          this.divisis = response.data;
        })
    },

    submitData() {
      if (this.password != this.confirmPassword) {
        this.message = "Password tidak cocok!";
        return
      }

      if (this.id_divisi == "" && this.id_divisi == "") {
        this.message = "Seluruh field harus terisi!"
        return
      }

      const data = {
        email: this.email,
        nama: this.nama,
        role: this.role,
        password: this.password,
        id_divisi: this.id_divisi,
        id_region: this.id_region,
      }

      axios.post("/registration", data)
      .then(() => {
        this.registered = true;
      })
    }
  }
}
</script>

<style scoped>
.none {
  display: none;
}

.true {
  display: block;
}
</style>