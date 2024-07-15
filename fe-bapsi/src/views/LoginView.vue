<template>
  <section class="full">
    <div class="wrapper">
      <h2>Login</h2>
      <span>
        <h3>Biro Administrasi, Perencanaan, dan Sistem Informasi</h3>
        <p>Universitas Gunadarma</p>
      </span>

      <form @submit.prevent="submitData">
        <span>
          <label for="email">Email</label>
          <input id="email" type="email" v-model="Email" placeholder="Email...">
        </span>
        <span>
          <label for="password">Password</label>
          <input id="password" type="password" v-model="Password" placeholder="Password...">
        </span>

        <span>
          <p class="red" v-if="message">{{ message }}</p>
        </span>

        <a href="#" @click="PushRoute('RegistrationView')">Belum memiliki akun?</a>

        <input type="submit" class="bt-1" value="Masuk">
      </form>

    </div>
  </section>
</template>

<script>
// Hallo, test 
import { axios } from "@/axios/config.js";

export default {
  data() {
    return {
      Email: "",
      Password: "",
      message: "",
    }
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

    submitData() {
      const data = {
        email: this.Email,
        password: this.Password,
      }
      axios.post("/login", data)
        .then((response) => {
          localStorage.setItem("tokenAuth", response.data.tokenJWT);
          window.location.reload();
        })
        .catch(() => {
          this.message = "Username atau password salah";
          setTimeout(() => {
            this.message = "";
          }, 3000)
        });
    }
  }
}
</script>