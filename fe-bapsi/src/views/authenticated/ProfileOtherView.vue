<!-- Halman untuk melihat profile orang lain -->
<!-- doc checked -->

<template>
  <section class="full">
    <div class="wrapper row">
      <div class="profile">
        <span class="garis">
          <h1>{{ profile.nama }}</h1>
          <h3>{{ profile.divisi }}</h3>
          <p>BAPSI {{ profile.region }}</p>
        </span>

        <span>
          <!-- Total akumulasi di bulan berjalan -->
          <h2>Akumulasi Juli 2024</h2>
          <span class="accumulation-data">
            <img src="@/assets/icons/clock.png" alt="time">
            <h3>{{ profile.akumulasi_jam }} Jam, {{ profile.akumulasi_menit }} Menit</h3>
          </span>

          <!-- hapus akun di sini -->
          <!-- hanya admin yang mendapat tampilah hapus user -->
          <span v-if="role == 'admin'">
            <br>
            <a href="#" class="red" @click="cofirmDeleteUser()">Hapus akun</a>
          </span>
        </span>

      </div>
      <div class="profile-image">
        <img src="@/assets/icons/blankProfile.jpg" alt="profile">
      </div>
    </div>
  </section>

  <section>
    <div class="wrapper">
      <span>
        <h1 class="title">Kehadiran</h1>
        <h3>Juli 2024</h3>
      </span>

      <p v-if="!kehadiran">Data masih kosong!!!</p>

      <div class="table" v-else>
        <table>
          <thead>
            <tr>
              <th>Tanggal</th>
              <th>Hari</th>
              <th>Jam Masuk</th>
              <th>Jam Pulang</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(data, index) in kehadiran" :key="index" :class="{ gray: index % 2 == 0 }">
              <td>{{ data.tanggal }}</td>
              <td>{{ data.hari }}</td>
              <td class="green">{{ data.masuk_at }}</td>
              <td class="red">{{ data.keluar_at }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </section>

  <!-- table daftar riwayat kehadiran harian di bulan berjalan -->
  <section v-if="inactive">
    <div class="wrapper">
      <span>
        <h1 class="title">Akumulasi</h1>
        <h3>2024</h3>
      </span>
      <div class="table">
        <table>
          <thead>
            <tr>
              <th>Tahun</th>
              <th>Bulan</th>
              <th>Akumulasi</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(data, index) in kehadiran" :key="index" :class="{ gray: index % 2 == 0 }">
              <td>{{ data.tanggal }}</td>
              <td>{{ data.hari }}</td>
              <td class="green">{{ data.masuk }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </section>

  <ConfirmationCard :nama="profile.nama" v-if="konfirmasiHapus" @kembali="handlerBatal()" @hapus="deleteUser()" />
</template>

<script>
</script>

<script>
import {axios} from "@/axios/config.js";
import ConfirmationCard from "@/components/alert/ConfirmationalCard.vue";

export default {
  data() {
    return {
      profile: Object,
      kehadiran: [],
      lengthKehadiran: 0,
      inactive: false,

      role: "",

      userID: "",
      konfirmasiHapus: false,
    }
  },

  components: {
    ConfirmationCard,
  },

  created() {
    this.userID = this.$route.params.id;
    this.getProfile();
    axios.get(`/profile/absen/${this.getYear()}/${this.getMonth()}/${this.userID}`)
    .then((response) => {
      this.kehadiran = response.data;
    });

    this.getRole();
  },

  methods: {
    cofirmDeleteUser() {
      this.konfirmasiHapus = true;
    },

    handlerBatal() {
      this.konfirmasiHapus = false;
    },

    getProfile() {
      axios.get(`/profile/${this.userID}`)
      .then((response) => {
        this.profile = response.data;
      })
      .catch((error) => {
      })
    },

    getRole() {
      axios.get("/role")
      .then((response) => {
        this.role = response.data.role;
      })
    },

    getMonth() {
      // Get current date
      const now = new Date();
      
      const mm = (now.getMonth() + 1).toString().padStart(2, '0');

      return mm;

      // const monthNames = [
      //   'January', 'February', 'March', 'April', 'May', 'June',
      //   'July', 'August', 'September', 'October', 'November', 'December'
      // ];

      // // Get current month name
      // const monthName = monthNames[now.getMonth()];
    },

    getYear() {
      const now = new Date();
      const yyyy = now.getFullYear();
      return yyyy;
    },

    deleteUser() {
      const data = {
        user_id: this.profile.id_user,
      }

      axios.delete(`/profile/${this.userID}`)
      .then((response) => {
        this.$router.push({name: 'DaftarStaff'});
      })
      .catch((error) => {
        console.log(error);
      })
    }
  }
}
</script>

<style scoped>
table {
  width: 600px;
  max-width: 90%;
  box-sizing: border-box;
  overflow: hidden;
  padding-top: 1em;
}

.table {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 100%;
  overflow: scroll;
  flex-direction: column;
}

section {
  box-sizing: border-box;
}

table tr {
  box-sizing: border-box;
}

table th,
table td {
  padding: 5px 10px;
  box-sizing: border-box;
  text-wrap: nowrap;
}

tr.gray {
  background-color: var(--gray);
}
</style>


<style scoped>
.profile h1 {
  font-size: 2.5rem;
  transition: 0.3s;
  font-weight: 600;
  color: var(--primary);
}

.profile h1:hover {
  color: var(--black);
}

.wrapper {
  justify-content: space-around;
}

.profile {
  display: flex;
  flex-direction: column;
  text-align: start;
  gap: 1rem;
}

.profile span.noData {
  gap: 1rem;
}

.wrapper {
  flex-wrap: wrap-reverse !important;
}

img {
  height: 500px;
  width: 400px;
  max-width: 100%;
  object-fit: cover;
  border-radius: 8px;
  transition: 0.3s;
}

span.accumulation-data {
  display: flex;
  gap: 1em;
  align-items: center;
}

.accumulation-data img {
  width: 28px;
  height: 28px;
}

@media screen and (max-width: 768px) {
  .profile h1 {
    font-size: 1.75rem;
  }

  img {
    height: 400px;
  }

  .profile {
    width: 100%;
  }

  .profile span.garis {
    border-left: 3px solid var(--primary);
    gap: 1rem;
    width: 100%;
    padding-left: 1rem;
    box-sizing: border-box;
  }

  span {
    width: 100%;
    text-align: left;
  }

  .table {
    padding: 0 1em;
    align-items: flex-start;
    justify-content: flex-start;
    flex-grow: initial;
    box-sizing: border-box;
  }
}
</style>