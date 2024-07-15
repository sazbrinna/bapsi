<template>
  <section class="full">
    <div class="wrapper">
      <h2>Registrasi</h2>
      <span>
        <h3>Biro Administrasi, Perencanaan, dan Sistem Informasi</h3>
        <p>Universitas Gunadarma</p>
      </span>

      <form @submit.prevent="SubmitData">

        <!-- First Section -->

        <div class="form-section" v-if="firstSection">
          <span>
            <label for="email">Email</label>
            <input id="email" type="email" v-model="Email" placeholder="Email..." required>
          </span>
          <span>
            <label for="password">Password</label>
            <input id="password" type="password" v-model="Password" placeholder="Password..." @blur="validatePassword"
              autocomplete="new-password" required>
          </span>
          <span>
            <label for="passwordconfirm">Konfirmasi Password</label>
            <input id="passwordconfirm" type="password" v-model="ConfirmationPassword"
              placeholder="Konfirmasi password..." @blur="validateConfirmPassword" autocomplete="new-Confirmpassword"
              required>
          </span>

          <div class="form-warning" v-if="passwordError">{{ passwordError }}</div>
          <div class="form-warning" v-if="confirmPasswordError">{{ confirmPasswordError }}</div>
          <div class="form-warning" v-if="status">{{ status }}</div>
          <div class="form-warning" v-if="emailNotFound">Email tidak terdaftar, silahkan hubungi admin atau staff</div>

          <a href="#" @click="PushRoute('LoginView')">Sudah memiliki akun?</a>

          <input type="button" class="bt-2" value="Selanjutnya" @click="NextSection()">
        </div>

        <!-- Second Section -->

        <div class="form-section" v-if="!firstSection">
          <span>
            <label for="">Nama</label>
            <input type="text" v-model="nama" placeholder="Masukkan nama kamu..." required>
          </span>
          <span>
            <label for="">Divisi</label>
            <h4>{{ divisi }}</h4>
          </span>
          <span>
            <label for="">Region</label>
            <h4>BAPSI {{ region }}</h4>
          </span>

          <!-- submit button -->
          <input type="submit" class="bt-1" value="Buat akun">
        </div>
      </form>
    </div>
  </section>
</template>

<script>
import { axios } from "@/axios/config.js";

export default {
  data() {
    return {
      Email: "",
      Password: "",
      ConfirmationPassword: "",
      StatusMessage: "",
      passwordError: "",
      confirmPasswordError: "",

      nama: "",
      divisi: "",
      region: "",

      firstSection: true,
      passwordReady: false,
      emailNotFound: false,
      status: "",
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

    NextSection() {
      if (!this.Email.length > 0 && !this.passwordReady) {
        this.status = "Kamu harus mengisi seluruh field"
        return;
      }

      axios.get(`/invitation/${this.Email}`)
        .then((response) => {
          this.divisi = response.data.divisi;
          this.region = response.data.region;
          this.firstSection = false;
        })
        .catch((error) => {
          if (error.response.status == "400") {
            this.emailNotFound = true;
            return;
          }
        })
    },

    SubmitData() {
      const data = {
        email: this.Email,
        password: this.Password,
        nama: this.nama,
        divisi: this.divisi,
        region: this.region,
      }

      axios.post("/registration", data)
        .then((response) => {
          localStorage.setItem("tokenAuth", response.data.tokenJWT);
          window.location.reload();
        })
        .catch(() => {
        });
    },

    validatePassword() {
      // Reset previous error message
      this.passwordError = '';

      // Check if password is at least 6 characters long
      if (this.Password.length < 6) {
        this.passwordError = 'Password must be at least 6 characters long.';
        return;
      }

      // Check if password contains at least one uppercase letter
      if (!/[A-Z]/.test(this.Password)) {
        this.passwordError = 'Password must contain at least one uppercase letter.';
        return;
      }

      // Check if password contains at least one non uppercase letter
      if (!/[a-z]/.test(this.Password)) {
        this.passwordError = 'Password must contain at least one non uppercase letter.';
        return;
      }

      // Check if password contains at least one digit
      if (!/\d/.test(this.Password)) {
        this.passwordError = 'Password must contain at least one digit.';
        return;
      }

      this.passwordReady = true;
    },

    validateConfirmPassword() {
      // Check if passwords match
      this.passwordError = '';

      if (this.Password !== this.ConfirmationPassword) {
        this.confirmPasswordError = 'Passwords do not match.';
      } else {
        this.confirmPasswordError = '';
      }
    }
  }
}
</script>