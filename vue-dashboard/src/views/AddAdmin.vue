<template>
  <div class="px-margin-mobile md:px-0 max-w-4xl mx-auto">
    <!-- Header Section -->
    <div class="mb-lg">
      <div class="flex items-center gap-sm mb-md">
        <router-link to="/admins" class="p-2 -ml-2 text-on-surface-variant transition-all duration-200 active:scale-95 hover:bg-surface-container-high rounded-lg">
          <span class="material-symbols-outlined">arrow_back</span>
        </router-link>
        <span class="font-label-caps text-label-caps text-on-surface-variant">BACK TO REGISTRY</span>
      </div>
      <h2 class="font-display-lg text-display-lg text-on-surface">{{ isAdminEdit ? 'Update Administrator' : 'Provision New Admin' }}</h2>
      <p class="text-on-surface-variant font-body-md mt-base" style="max-width: 36rem;">
        {{ isAdminEdit ? 'Modify system credentials and update operational permissions for this administrator.' : 'Assign system credentials and define operational permissions for the new technical administrator.' }}
        All actions are logged under the master security protocol.
      </p>
    </div>

    <!-- Form Content -->
    <form class="space-y-gutter" @submit.prevent="submitForm">
      <!-- Avatar Selection -->
      <section class="glass-panel p-md rounded-xl">
        <div class="flex flex-col md:flex-row items-center gap-lg">
          <div class="relative group">
            <div class="w-24 h-24 rounded-full bg-surface-container-highest border-2 border-dashed border-outline-variant flex items-center justify-center overflow-hidden transition-all group-hover:border-primary">
              <img v-if="avatarPreview" :src="avatarPreview" class="w-full h-full object-cover" alt="Avatar preview"/>
              <span v-else class="material-symbols-outlined text-4xl text-on-surface-variant">add_a_photo</span>
            </div>
            <input accept="image/*" class="hidden" ref="avatarInput" type="file" @change="handleAvatarUpload"/>
            <button class="absolute -bottom-1 -right-1 bg-primary text-on-primary w-8 h-8 rounded-full flex items-center justify-center shadow-lg transition-transform hover:scale-110 active:scale-95" type="button" @click="$refs.avatarInput.click()">
              <span class="material-symbols-outlined text-sm">edit</span>
            </button>
          </div>
          <div class="flex-1 space-y-base text-center md:text-left">
            <p class="font-title-sm text-title-sm text-on-surface">Administrative Identity</p>
            <p class="text-on-surface-variant text-body-md">Upload a 1:1 aspect ratio profile image. This will be used for authentication logs and terminal headers.</p>
          </div>
        </div>
      </section>

      <!-- Identity Fields -->
      <section class="glass-panel p-md rounded-xl grid grid-cols-1 md:grid-cols-2 gap-gutter">
        <div class="space-y-xs col-span-1">
          <label class="font-label-caps text-label-caps text-on-surface-variant block uppercase tracking-wider">Full Name</label>
          <div class="relative">
            <input v-model="form.name" class="w-full bg-surface-container-lowest border border-outline-variant rounded-lg px-md py-3 text-on-surface font-data-md transition-all focus:bg-surface-container focus:border-primary focus:ring-1 focus:ring-primary outline-none" placeholder="e.g., Marcus Thorne" type="text"/>
            <span class="material-symbols-outlined absolute right-3 top-3 text-on-surface-variant text-sm">badge</span>
          </div>
        </div>
        <div class="space-y-xs col-span-1">
          <label class="font-label-caps text-label-caps text-on-surface-variant block uppercase tracking-wider">Work Email</label>
          <div class="relative">
            <input v-model="form.email" class="w-full bg-surface-container-lowest border border-outline-variant rounded-lg px-md py-3 text-on-surface font-data-md transition-all focus:bg-surface-container focus:border-primary focus:ring-1 focus:ring-primary outline-none" placeholder="m.thorne@industrial-node.io" type="email"/>
            <span class="material-symbols-outlined absolute right-3 top-3 text-on-surface-variant text-sm">mail</span>
          </div>
        </div>
        <div class="space-y-xs col-span-1 md:col-span-2">
          <label class="font-label-caps text-label-caps text-on-surface-variant block uppercase tracking-wider">Encrypted Phone Line</label>
          <div class="relative">
            <input v-model="form.phone" class="w-full bg-surface-container-lowest border border-outline-variant rounded-lg px-md py-3 text-on-surface font-data-md transition-all focus:bg-surface-container focus:border-primary focus:ring-1 focus:ring-primary outline-none" placeholder="+1 (555) 902-1142" type="tel"/>
            <span class="material-symbols-outlined absolute right-3 top-3 text-on-surface-variant text-sm">dock_to_bottom</span>
          </div>
        </div>
      </section>

      <!-- Role Selection -->
      <section class="glass-panel p-md rounded-xl space-y-md">
        <label class="font-label-caps text-label-caps text-on-surface-variant block uppercase tracking-wider">Access Tier Assignment</label>
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-sm">
          <label v-for="role in roles" :key="role.value" class="cursor-pointer group">
            <input class="peer hidden" name="roles" type="checkbox" :value="role.value" v-model="form.roles"/>
            <div class="h-full p-md bg-surface-container-low border border-outline-variant rounded-lg flex flex-col items-center text-center transition-all peer-checked:border-primary peer-checked:bg-primary/5 hover:bg-surface-container-high group-active:scale-95">
              <span class="material-symbols-outlined text-on-surface-variant mb-sm" style="font-variation-settings: 'FILL' 1;">{{ role.icon }}</span>
              <p class="font-label-caps text-label-caps text-on-surface mb-xs">{{ role.label }}</p>
              <p class="text-[10px] text-on-surface-variant leading-tight">{{ role.description }}</p>
            </div>
          </label>
        </div>
      </section>

      <!-- Direct Permissions -->
      <section class="glass-panel p-md rounded-xl space-y-md">
        <label class="font-label-caps text-label-caps text-on-surface-variant block uppercase tracking-wider">Direct Security Permissions</label>
        <p class="text-[11px] text-on-surface-variant mb-md">Directly assign granular permissions that override or augment role-based access.</p>
        <div class="grid grid-cols-1 sm:grid-cols-2 gap-sm">
          <label v-for="perm in availablePermissions" :key="perm.value" class="flex items-center gap-sm p-sm bg-surface-container-low border border-white/5 rounded-lg cursor-pointer hover:bg-white/5 transition-colors">
            <input type="checkbox" :value="perm.value" v-model="form.permissions" class="w-4 h-4 rounded border-outline-variant text-primary focus:ring-primary bg-transparent"/>
            <div class="flex flex-col">
              <span class="font-title-sm text-[12px] text-on-surface">{{ perm.label }}</span>
              <span class="text-[10px] text-on-surface-variant">{{ perm.description }}</span>
            </div>
          </label>
        </div>
      </section>

      <!-- Action Buttons -->
      <div class="flex flex-col sm:flex-row gap-md pt-lg">
        <button class="flex-1 bg-primary text-on-primary font-title-sm text-title-sm py-4 rounded-xl flex items-center justify-center gap-sm transition-all duration-200 hover:brightness-110 active:scale-[0.98] shadow-lg shadow-primary/10" type="submit" :disabled="isSubmitting">
          <span v-if="isSubmitting" class="material-symbols-outlined animate-spin">refresh</span>
          <span v-else-if="isSuccess" class="material-symbols-outlined">check_circle</span>
          <span v-else class="material-symbols-outlined">save</span>
          {{ submitLabel }}
        </button>
        <router-link to="/admins" class="px-xl border border-outline-variant text-on-surface-variant font-title-sm text-title-sm py-4 rounded-xl transition-all hover:bg-surface-container-high active:scale-95 text-center">
          Cancel
        </router-link>
      </div>
    </form>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { adminApi } from '../api/admins'
import { roleApi } from '../api/roles'

const router = useRouter()
const route = useRoute()

const avatarFile = ref(null)
const avatarPreview = ref(null)
const isSubmitting = ref(false)
const isSuccess = ref(false)
const roles = ref([])
const availablePermissions = ref([
    { value: 'view-dashboard', label: 'View Dashboard', description: 'Access to system overview.' },
    { value: 'manage-admins', label: 'Manage Admins', description: 'Create/Edit/Delete personnel.' },
    { value: 'manage-roles', label: 'Manage Roles', description: 'Configure access tiers.' },
    { value: 'manage-iot-devices', label: 'Manage IoT', description: 'Control field hardware.' },
    { value: 'view-system-logs', label: 'View Logs', description: 'Audit security events.' }
])
const isAdminEdit = computed(() => !!route.params.id)

const form = ref({
  name: '',
  email: '',
  phone: '',
  password: 'password123', // Default for new admins as per Postman
  roles: [],
  permissions: []
})

const fetchRoles = async () => {
    try {
        const data = await roleApi.getRoles()
        roles.value = data.map(r => ({
            value: r.Name.toLowerCase(),
            label: r.Name.toUpperCase(),
            icon: r.Name.toLowerCase() === 'root' ? 'terminal' : 
                  r.Name.toLowerCase() === 'security' ? 'security' : 
                  r.Name.toLowerCase() === 'network' ? 'hub' : 'visibility',
            description: r.Description || 'Standard access role.'
        }))
        if (roles.value.length > 0 && form.value.roles.length === 0) {
            form.value.roles = [roles.value[0].value]
        }
    } catch (err) {
        console.error('Failed to fetch roles:', err)
        // Fallback roles if API fails
        roles.value = [
            { value: 'root', label: 'ROOT', icon: 'terminal', description: 'Full kernel access & system overrides.' },
            { value: 'security', label: 'SECURITY', icon: 'security', description: 'Firewall mgmt & threat mitigation.' }
        ]
    }
}

const fetchAdminData = async () => {
    if (!isAdminEdit.value) return
    try {
        const admin = await adminApi.getAdmin(route.params.id)
        // Flexible mapping to handle both PascalCase (Go) and lowercase/camelCase
        form.value = {
            name: admin.Name || admin.name || '',
            email: admin.Email || admin.email || '',
            phone: admin.Phone || admin.phone || '',
            password: '', // Don't pre-fill password for security
            roles: (admin.roles || admin.Roles || []).map(r => (r.Name || r.name).toLowerCase()),
            permissions: (admin.permissions || admin.Permissions || []).map(p => (p.Name || p.name).toLowerCase())
        }
        // Load existing avatar from backend
        const existingAvatar = admin.Avatar || admin.avatar
        if (existingAvatar) {
            avatarPreview.value = existingAvatar.startsWith('http') ? existingAvatar
                : existingAvatar.startsWith('/') ? existingAvatar
                : `/${existingAvatar}`
        }
    } catch (err) {
        console.error('Failed to fetch admin data:', err)
    }
}

const submitLabel = computed(() => {
  if (isSuccess.value) return isAdminEdit.value ? 'Admin Updated' : 'Admin Provisioned'
  if (isSubmitting.value) return 'Processing...'
  return isAdminEdit.value ? 'Update Admin' : 'Save Admin'
})

const handleAvatarUpload = (e) => {
  const file = e.target.files[0]
  if (file) {
    avatarFile.value = file // Store the actual file
    const reader = new FileReader()
    reader.onload = (ev) => {
      avatarPreview.value = ev.target.result
    }
    reader.readAsDataURL(file)
  }
}

const submitForm = async () => {
    isSubmitting.value = true
    try {
        const adminData = new FormData()
        adminData.append('name', form.value.name)
        adminData.append('email', form.value.email)
        adminData.append('phone', form.value.phone)
        
        // Append multiple roles
        form.value.roles.forEach(role => {
            adminData.append('roles', role)
        })

        // Append multiple permissions
        form.value.permissions.forEach(perm => {
            adminData.append('permissions', perm)
        })

        if (!isAdminEdit.value) {
            adminData.append('password', form.value.password)
        }

        if (avatarFile.value) {
            adminData.append('avatar', avatarFile.value)
        }

        if (isAdminEdit.value) {
            await adminApi.updateAdmin(route.params.id, adminData)
        } else {
            await adminApi.createAdmin(adminData)
        }

        isSuccess.value = true
        setTimeout(() => {
            router.push('/admins')
        }, 1500)
    } catch (err) {
        console.error('Submission failed:', err)
        alert('Failed to save administrator record.')
    } finally {
        isSubmitting.value = false
    }
}

onMounted(() => {
    fetchRoles()
    fetchAdminData()
})
</script>

<style scoped>
.glass-panel {
    background: rgba(23, 31, 51, 0.7);
    backdrop-filter: blur(12px);
    border: 1px solid rgba(255, 255, 255, 0.08);
}
</style>
