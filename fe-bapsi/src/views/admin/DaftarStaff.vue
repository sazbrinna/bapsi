<template>
  <TitlePart title="Daftar Asisten" description="Juli 2024" />

  <section>
    <div class="wrapper">
      <p class="pointer"></p>
      <!-- <p class="link" @click="PushRoute('InvitationView')">Undang staff baru</p> -->
      <div class="table">
        <table>
          <thead>
            <tr>
              <th class="primary">No</th>
              <th class="primary">Nama</th>
              <th class="primary">Region</th>
              <th class="primary">Divisi</th>
              <th class="primary">Akumulasi Kehadiran</th>
            </tr>
          </thead>

          <tbody>
            <tr v-for="(staff, index) in staffs" :key="index" :class="{ gray: index % 2 == 0 }">
              <td>{{ index + 1 }}</td>
              <td class="link" @click.prevent="handlerProfile(staff.id_user)"><strong>{{ staff.nama }}</strong></td>
              <td>{{ staff.region }}</td>
              <td>{{ staff.divisi }}</td>
              <td class="green">{{ staff.tot_jam }} jam, {{ staff.tot_menit }} menit</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </section>

  <section>
    <div class="wrapper">
      <RequestLogin :requests="requests" />
    </div>
  </section>
</template>

<script>
import { axios } from "@/axios/config.js";
import TitlePart from "@/components/TitlePart.vue";

import RequestLogin from "@/components/admin/WrapRequest.vue";

export default {
  components: {
    TitlePart,
    RequestLogin
  },

  data() {
    return {
      staffs: [],
      requests: [],

      idInvitation: "",
    }
  },

  created() {
    this.getRequest();
    this.getStaff();
  },

  methods: {
    getStaff() {
      axios.get("/staffs")
        .then((response) => {
          this.staffs = response.data.sort((a, b) => a.nama.localeCompare(b.nama));
        })
    },

    getRequest() {
      axios.get("/request")
        .then((response) => {
          this.requests = response.data;
        })
        .catch((error) => {
          console.log(error);
        })
    },

    handlerProfile(path) {
      window.scrollTo({
        top: 0,
        behavior: 'smooth'
      });

      this.$router.push({ path: "/profile/" + path });
    },

    PushRoute(pageName) {
      window.scrollTo({
        top: 0,
        behavior: 'smooth'
      });

      this.$router.push({ name: pageName })
    },
  }
}
</script>

<style scoped>
td.link {
  color: var(--black) !important;
}
</style>