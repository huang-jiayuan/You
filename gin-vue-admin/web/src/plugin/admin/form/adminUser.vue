
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="用户的ID:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="管理员的密码:" prop="adminPassword">
          <el-input v-model="formData.adminPassword" :clearable="true"  placeholder="请输入管理员的密码" />
       </el-form-item>
        <el-form-item label="管理员账号:" prop="adminName">
          <el-input v-model="formData.adminName" :clearable="true"  placeholder="请输入管理员账号" />
       </el-form-item>
        <el-form-item label="用户的账号状态(0-封禁,1-冻结,2-禁言):" prop="status">
           <el-select v-model="formData.status" placeholder="请选择用户的账号状态(0-封禁,1-冻结,2-禁言)" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in userstatusOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="管理员等级（1普通,2超级）:" prop="idAdmin">
           <el-select v-model="formData.idAdmin" placeholder="请选择管理员等级（1普通,2超级）" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in 管理员Options" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createAdminUser,
  updateAdminUser,
  findAdminUser
} from '@/plugin/admin/api/adminUser'

defineOptions({
    name: 'AdminUserForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const userstatusOptions = ref([])
const 管理员Options = ref([])
const formData = ref({
            userId: undefined,
            adminPassword: '',
            adminName: '',
            status: '',
            idAdmin: '',
        })
// 验证规则
const rule = reactive({
               userId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               adminPassword : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               adminName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               status : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               idAdmin : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findAdminUser({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    userstatusOptions.value = await getDictFunc('userstatus')
    管理员Options.value = await getDictFunc('管理员')
}

init()
// 保存按钮
const save = async() => {
      btnLoading.value = true
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return btnLoading.value = false
            let res
           switch (type.value) {
             case 'create':
               res = await createAdminUser(formData.value)
               break
             case 'update':
               res = await updateAdminUser(formData.value)
               break
             default:
               res = await createAdminUser(formData.value)
               break
           }
           btnLoading.value = false
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
