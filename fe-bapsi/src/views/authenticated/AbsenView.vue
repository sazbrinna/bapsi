<template>
  <div class="container">
    <div class="wrapper g2">
      <TitlePart title="Absensi Kehadiran Loket" description="Januari - 2024" />

      <button class="bt-1">Absen Pulang</button>

      <div class="row g1">
        <select v-model="selectedYear">
          <option v-for="year in years" :key="year" :value="year">{{ year }}</option>
        </select>

        <select v-model="selectedMonth">
          <option v-for="(month, index) in months" :key="index" :value="index + 1">{{ month }}</option>
        </select>
      </div>

      <table>
        <tr>
          <th>Tanggal</th>
          <th>Jam Masuk</th>
          <th>Jam Pulang</th>
        </tr>
        <tr>
          <td>13-01-2024</td>
          <td>12.20</td>
          <td>14.00</td>
        </tr>
      </table>
    </div>
  </div>
</template>

<script>
import TitlePart from "@/components/TitlePart.vue";

export default {
  components: {
    TitlePart
  },

  data() {
    return {
      selectedYear: new Date().getFullYear(),
      selectedMonth: new Date().getMonth() + 1,
      years: [],
      months: [
        'January', 'February', 'March', 'April', 'May', 'June',
        'July', 'August', 'September', 'October', 'November', 'December'
      ]
    };
  },

  mounted() {
    // Populate years dropdown with a range of years
    const currentYear = new Date().getFullYear();
    for (let year = currentYear - 10; year <= currentYear + 10; year++) {
      this.years.push(year);
    }
  }
};
</script>