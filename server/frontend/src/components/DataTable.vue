<template>
  <div class="data-table">
    <div class="filters">
      <input v-model="search" placeholder="搜索..." />
      <select v-model="sortBy">
        <option value="id">ID</option>
        <option value="connection_desc">连接描述</option>
        <option value="protocol">协议</option>
        <option value="total_time_ms">总时间</option>
        <option value="request_size">请求大小</option>
        <option value="response_size">响应大小</option>
        <option value="process">进程</option>
        <option value="net_internal_time_ms">内部时间</option>
        <option value="read_socket_time_ms">读取时间</option>
        <option value="start_time">开始时间</option>
        <option value="request">请求</option>
        <option value="response">响应</option>
      </select>
      <button @click="sortOrder = sortOrder === 'asc' ? 'desc' : 'asc'">
        {{ sortOrder === 'asc' ? '升序' : '降序' }}
      </button>
    </div>

    <table>
      <thead>
        <tr>
          <th>ID</th>
          <th>连接描述</th>
          <th>协议</th>
          <th>总时间(ms)</th>
          <th>请求大小</th>
          <th>响应大小</th>
          <th>进程</th>
          <th>内部时间(ms)</th>
          <th>读取时间(ms)</th>
          <th>开始时间</th>
          <th>请求</th>
          <th>响应</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="record in sortedAndFilteredRecords" :key="record.id">
          <td>{{ record.id }}</td>
          <td>{{ record.connection_desc }}</td>
          <td>{{ record.protocol }}</td>
          <td>{{ record.total_time_ms }}</td>
          <td>{{ record.request_size }}</td>
          <td>{{ record.response_size }}</td>
          <td>{{ record.process }}</td>
          <td>{{ record.net_internal_time_ms }}</td>
          <td>{{ record.read_socket_time_ms }}</td>
          <td>{{ formatDate(record.start_time) }}</td>
          <td>{{ record.request }}</td>
          <td>{{ record.response }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  data() {
    return {
      records: [],
      search: '',
      sortBy: 'id',
      sortOrder: 'asc'
    }
  },
  computed: {
    sortedAndFilteredRecords() {
      return this.records
        .filter(record => 
          Object.values(record).some(value => 
            String(value).toLowerCase().includes(this.search.toLowerCase())
          )
        )
        .sort((a, b) => {
          const modifier = this.sortOrder === 'asc' ? 1 : -1
          if (typeof a[this.sortBy] === 'string') {
            return a[this.sortBy].localeCompare(b[this.sortBy]) * modifier
          } else if (typeof a[this.sortBy] === 'number') {
            return (a[this.sortBy] - b[this.sortBy]) * modifier
          } else if (a[this.sortBy] instanceof Date) {
            return (new Date(a[this.sortBy]) - new Date(b[this.sortBy])) * modifier
          }
          return 0
        })
    }
  },
  methods: {
    async fetchData() {
      try {
        const response = await fetch('/api/records')
        this.records = await response.json()
      } catch (error) {
        console.error('Error fetching data:', error)
      }
    },
    formatDate(date) {
      return new Date(date).toLocaleString()
    }
  },
  mounted() {
    this.fetchData()
    // 每10秒刷新一次数据
    setInterval(this.fetchData, 10000)
  }
}
</script>

<style scoped>
.data-table {
  padding: 20px;
}

.filters {
  margin-bottom: 20px;
}

table {
  width: 100%;
  border-collapse: collapse;
}

th, td {
  border: 1px solid #ddd;
  padding: 8px;
  text-align: left;
}

th {
  background-color: #f2f2f2;
}

tr:nth-child(even) {
  background-color: #f9f9f9;
}
</style>