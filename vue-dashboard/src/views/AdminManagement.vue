<template>
  <div class="space-y-lg">
    <div class="flex flex-col sm:flex-row justify-between items-start sm:items-end gap-4">
      <div>
        <h1 class="font-headline-md text-headline-md text-on-surface">Admin Management</h1>
        <p class="font-body-md text-body-md text-on-surface-variant mt-xs">Manage system access levels and personnel security.</p>
      </div>
      <router-link to="/admins/add" class="bg-primary hover:bg-primary-fixed-dim text-on-primary font-bold px-lg py-sm rounded-lg transition-all duration-200 flex items-center gap-sm shadow-sm whitespace-nowrap">
        <span class="material-symbols-outlined">person_add</span>
        <span>Add Admin</span>
      </router-link>
    </div>

    <!-- Stats Grid -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-md">
      <div class="glass-panel p-md rounded-xl col-span-1">
        <p class="font-label-caps text-label-caps text-on-surface-variant">Total Admins</p>
        <div class="flex items-end justify-between mt-sm">
          <span class="font-display-lg text-display-lg text-on-surface">{{ admins.length }}</span>
          <span class="text-primary font-data-md text-data-md">+{{ Math.floor(admins.length * 0.1) }} this month</span>
        </div>
      </div>
      <div class="glass-panel p-md rounded-xl col-span-1">
        <p class="font-label-caps text-label-caps text-on-surface-variant">Active Now</p>
        <div class="flex items-end justify-between mt-sm">
          <span class="font-display-lg text-display-lg text-primary">{{ String(Math.ceil(admins.length * 0.25)).padStart(2, '0') }}</span>
          <div class="flex -space-x-2">
            <div class="w-6 h-6 rounded-full border-2 border-surface-container overflow-hidden">
              <img class="w-full h-full object-cover" src="../assets/images/user_engineer.jpg"/>
            </div>
            <div class="w-6 h-6 rounded-full border-2 border-surface-container overflow-hidden">
              <img class="w-full h-full object-cover" src="../assets/images/user_ito.jpg"/>
            </div>
          </div>
        </div>
      </div>
      <div class="glass-panel p-md rounded-xl col-span-1 md:col-span-2 overflow-hidden relative">
        <div class="absolute inset-0 bg-gradient-to-r from-primary/5 to-transparent"></div>
        <p class="font-label-caps text-label-caps text-on-surface-variant relative z-10">Security Health</p>
        <div class="flex items-center gap-lg mt-sm relative z-10">
          <div class="flex-1">
            <div class="h-2 w-full bg-surface-container-highest rounded-full overflow-hidden">
              <div class="h-full bg-primary w-[92%]"></div>
            </div>
            <p class="font-data-md text-data-md mt-xs">92% Compliance</p>
          </div>
          <span class="material-symbols-outlined text-primary text-4xl">verified_user</span>
        </div>
      </div>
    </div>

    <!-- Admin Table -->
    <div class="glass-panel rounded-xl overflow-hidden">
      <div class="px-md py-md bg-white/5 border-b border-white/10 flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4">
        <div class="flex flex-col sm:flex-row items-start sm:items-center gap-md w-full sm:w-auto">
          <span class="font-title-sm text-title-sm">Registry</span>
          <div class="flex items-center bg-surface-container-lowest border border-white/10 px-sm py-1 rounded w-full sm:w-auto">
            <span class="material-symbols-outlined text-[18px] text-on-surface-variant mr-xs">search</span>
            <input v-model="searchQuery" class="bg-transparent border-none focus:ring-0 font-body-md text-body-md text-on-surface py-0 w-full sm:w-48 outline-none" placeholder="Filter personnel..." type="text"/>
          </div>
        </div>
        <div class="flex gap-sm w-full sm:w-auto justify-end">
          <button class="p-2 hover:bg-surface-variant rounded-lg transition-colors text-on-surface-variant">
            <span class="material-symbols-outlined">filter_list</span>
          </button>
          <button class="p-2 hover:bg-surface-variant rounded-lg transition-colors text-on-surface-variant">
            <span class="material-symbols-outlined">download</span>
          </button>
        </div>
      </div>
      <div class="overflow-x-auto custom-scrollbar">
        <table class="w-full border-collapse min-w-[600px]">
          <thead>
            <tr class="bg-surface-container-high/50 text-left border-b border-white/10">
              <th class="px-md py-3 font-label-caps text-label-caps text-on-surface-variant uppercase">Administrator</th>
              <th class="px-md py-3 font-label-caps text-label-caps text-on-surface-variant uppercase">Email Address</th>
              <th class="px-md py-3 font-label-caps text-label-caps text-on-surface-variant uppercase">Access Role</th>
              <th class="px-md py-3 font-label-caps text-label-caps text-on-surface-variant uppercase">Last Login</th>
              <th class="px-md py-3 font-label-caps text-label-caps text-on-surface-variant uppercase text-right">Actions</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-white/5">
            <tr v-if="isLoading">
              <td colspan="5" class="px-md py-12 text-center">
                <span class="material-symbols-outlined animate-spin text-primary text-3xl">refresh</span>
                <p class="font-label-caps text-on-surface-variant mt-sm">Decrypting Personnel Records...</p>
              </td>
            </tr>
            <tr v-else-if="error">
              <td colspan="5" class="px-md py-12 text-center">
                <span class="material-symbols-outlined text-error text-3xl">report_problem</span>
                <p class="font-label-caps text-error mt-sm">{{ error }}</p>
                <button @click="fetchAdmins" class="mt-md text-primary font-title-sm hover:underline">Retry Connection</button>
              </td>
            </tr>
            <tr v-else-if="filteredAdmins.length === 0">
              <td colspan="5" class="px-md py-12 text-center text-on-surface-variant font-body-md">
                {{ admins.length === 0 ? 'No personnel records found in the registry.' : 'No personnel match your search criteria.' }}
              </td>
            </tr>
            <tr v-for="admin in filteredAdmins" :key="admin.id || admin.email" class="hover:bg-white/5 transition-colors">
              <td class="px-md py-3">
                <div class="flex items-center gap-sm">
                  <div v-if="admin.image" class="w-8 h-8 rounded-full overflow-hidden border border-white/10">
                    <img :src="admin.image" class="w-full h-full object-cover"/>
                  </div>
                  <div v-else class="w-8 h-8 rounded-full flex items-center justify-center font-bold text-xs" :class="admin.bgClass">
                    {{ admin.initials }}
                  </div>
                  <span class="font-data-md text-data-md">{{ admin.name }}</span>
                </div>
              </td>
              <td class="px-md py-3 font-data-md text-data-md text-on-surface-variant">{{ admin.email }}</td>
              <td class="px-md py-3">
                <div class="flex flex-wrap gap-xs">
                  <span v-for="role in admin.roles" :key="role" class="px-2 py-0.5 rounded border font-label-caps text-[10px] uppercase" :class="admin.roleClass">
                    {{ role }}
                  </span>
                  <span v-if="!admin.roles || admin.roles.length === 0" class="px-2 py-0.5 rounded border font-label-caps text-[10px] uppercase bg-surface-container text-on-surface-variant border-outline-variant">
                    No Role
                  </span>
                </div>
              </td>
              <td class="px-md py-3 font-data-md text-data-md text-on-surface-variant">{{ admin.lastLogin }}</td>
              <td class="px-md py-3 text-right">
                <div class="flex justify-end gap-xs">
                  <router-link :to="`/admins/edit/${admin.id}`" class="p-1.5 hover:bg-primary/20 hover:text-primary rounded-lg transition-all text-on-surface-variant active:scale-90" title="Edit Personnel">
                    <span class="material-symbols-outlined text-[20px]">edit</span>
                  </router-link>
                  <button @click="deleteAdmin(admin.id)" class="p-1.5 hover:bg-error-container/20 hover:text-error rounded-lg transition-all text-on-surface-variant active:scale-90" title="Revoke Access">
                    <span class="material-symbols-outlined text-[20px]">delete</span>
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- Pagination Controls -->
    <div class="mt-md flex items-center justify-between bg-surface-container-low/30 p-md rounded-xl border border-white/5">
      <div class="text-on-surface-variant font-label-caps text-[11px]">
        Showing {{ (currentPage - 1) * 10 + 1 }} to {{ Math.min(currentPage * 10, totalItems) }} of {{ totalItems }} records
      </div>
      <div class="flex items-center gap-xs">
        <button @click="changePage(currentPage - 1)" :disabled="currentPage === 1" class="p-2 hover:bg-white/5 rounded-lg disabled:opacity-30 disabled:hover:bg-transparent transition-all">
          <span class="material-symbols-outlined text-[20px]">chevron_left</span>
        </button>
        <div class="flex items-center gap-xs">
          <button v-for="page in totalPages" :key="page" @click="changePage(page)" class="w-8 h-8 rounded-lg font-data-md text-[12px] transition-all" :class="currentPage === page ? 'bg-primary text-on-primary' : 'hover:bg-white/5 text-on-surface-variant'">
            {{ page }}
          </button>
        </div>
        <button @click="changePage(currentPage + 1)" :disabled="currentPage === totalPages" class="p-2 hover:bg-white/5 rounded-lg disabled:opacity-30 disabled:hover:bg-transparent transition-all">
          <span class="material-symbols-outlined text-[20px]">chevron_right</span>
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { adminApi } from '../api/admins'

const admins = ref([])
const searchQuery = ref('')
const isLoading = ref(false)
const error = ref('')

// Pagination
const currentPage = ref(1)
const totalItems = ref(0)
const totalPages = computed(() => Math.ceil(totalItems.value / 10))

const changePage = (page) => {
  if (page < 1 || page > totalPages.value) return
  currentPage.value = page
  fetchAdmins()
}

const filteredAdmins = computed(() => {
  if (!searchQuery.value) return admins.value
  const query = searchQuery.value.toLowerCase()
  return admins.value.filter(admin => 
    admin.name.toLowerCase().includes(query) || 
    admin.email.toLowerCase().includes(query) ||
    admin.roles.some(role => role.toLowerCase().includes(query))
  )
})

const resolveAvatar = (path) => {
  if (!path) return null
  // Already an absolute URL (http/https) — use as-is
  if (path.startsWith('http://') || path.startsWith('https://')) return path
  // Already starts with / — use as-is
  if (path.startsWith('/')) return path
  // Relative backend path like "uploads/avatars/123.jpg" → "/uploads/avatars/123.jpg"
  return `/${path}`
}

const fetchAdmins = async () => {
  isLoading.value = true
  error.value = ''
  try {
    const response = await adminApi.getAdmins(currentPage.value, 10)
    totalItems.value = response.total
    const data = response.data
    // Map API data (Support both PascalCase and lowercase) to UI format
    admins.value = data.map(admin => {
      const name = admin.Name || admin.name || 'Admin';
      const roleList = admin.roles || admin.Roles || [];
      return {
        id: admin.ID || admin.id,
        name: name,
        email: admin.Email || admin.email,
        roles: roleList.map(r => r.Name || r.name),
        image: resolveAvatar(admin.Avatar || admin.avatar),
        initials: name.split(' ').map(n => n[0]).join('').toUpperCase().substring(0, 2),
        bgClass: 'bg-primary/20 text-primary',
        roleClass: 'bg-primary/10 text-primary border-primary/20',
        lastLogin: 'Active Now'
      }
    })
  } catch (err) {
    error.value = 'Failed to load personnel registry.'
    console.error(err)
  } finally {
    isLoading.value = false
  }
}

const deleteAdmin = async (id) => {
  if (!confirm('Are you sure you want to deauthorize this administrator?')) return
  
  try {
    await adminApi.deleteAdmin(id)
    admins.value = admins.value.filter(a => a.id !== id)
  } catch (err) {
    alert('Security override failed: Could not delete admin.')
    console.error(err)
  }
}

onMounted(() => {
  fetchAdmins()
})
</script>
