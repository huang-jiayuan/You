
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="登录账号(手机号/第三方账号):" prop="username">
    <el-input v-model="formData.username" :clearable="true" placeholder="请输入登录账号(手机号/第三方账号)" />
</el-form-item>
        <el-form-item label="登录密码(SHA-256加密):" prop="password">
    <el-input v-model="formData.password" :clearable="true" placeholder="请输入登录密码(SHA-256加密)" />
</el-form-item>
        <el-form-item label="用户昵称:" prop="nickname">
    <el-input v-model="formData.nickname" :clearable="true" placeholder="请输入用户昵称" />
</el-form-item>
        <el-form-item label="头像URL:" prop="avatar">
    <el-input v-model="formData.avatar" :clearable="true" placeholder="请输入头像URL" />
</el-form-item>
        <el-form-item label="性别(0-保密,1-男,2-女):" prop="gender">
    <el-select v-model="formData.gender" placeholder="请选择性别(0-保密,1-男,2-女)" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in genderOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="声音标签(逗号分隔):" prop="voiceTag">
    <el-select v-model="formData.voiceTag" placeholder="请选择声音标签(逗号分隔)" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in voiceOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="兴趣标签(逗号分隔):" prop="interestTags">
    <el-select v-model="formData.interestTags" placeholder="请选择兴趣标签(逗号分隔)" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in interesttagsOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="用户等级:" prop="level">
    <el-input v-model.number="formData.level" :clearable="true" placeholder="请输入用户等级" />
</el-form-item>
        <el-form-item label="会员状态(0-非会员,1-月会员,2-年会员):" prop="vipStatus">
    <el-select v-model="formData.vipStatus" placeholder="请选择会员状态(0-非会员,1-月会员,2-年会员)" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in vipstatusOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="会员过期时间:" prop="vipExpireTime">
    <el-date-picker v-model="formData.vipExpireTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="虚拟币余额:" prop="balance">
    <el-input-number v-model="formData.balance" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
        <el-form-item label="钻石数量:" prop="diamond">
    <el-input v-model.number="formData.diamond" :clearable="true" placeholder="请输入钻石数量" />
</el-form-item>
        <el-form-item label="注册时间:" prop="registerTime">
    <el-date-picker v-model="formData.registerTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="最后登录时间:" prop="lastLoginTime">
    <el-date-picker v-model="formData.lastLoginTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="最后登录IP:" prop="lastLoginIp">
    <el-input v-model="formData.lastLoginIp" :clearable="true" placeholder="请输入最后登录IP" />
</el-form-item>
        <el-form-item label="账号状态(0-正常,1-封禁,2-冻结,3-禁言):" prop="status">
    <el-select v-model="formData.status" placeholder="请选择账号状态(0-正常,1-封禁,2-冻结,3-禁言)" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in userstatusOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="冻结/封禁原因:" prop="freezeReason">
    <el-input v-model="formData.freezeReason" :clearable="true" placeholder="请输入冻结/封禁原因" />
</el-form-item>
        <el-form-item label="冻结/封禁结束时间:" prop="freezeEndTime">
    <el-date-picker v-model="formData.freezeEndTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="实名认证状态（0：未认证1：已认证2）:" prop="isAuth">
    <el-select v-model="formData.isAuth" placeholder="请选择实名认证状态（0：未认证1：已认证2）" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in is_authOptions" :key="key" :label="item.label" :value="item.value" />
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
  createUser,
  updateUser,
  findUser
} from '@/api/users/user'

defineOptions({
    name: 'UserForm'
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
const genderOptions = ref([])
const voiceOptions = ref([])
const interesttagsOptions = ref([])
const userstatusOptions = ref([])
const vipstatusOptions = ref([])
const is_authOptions = ref([])
const formData = ref({
            username: '',
            password: '',
            nickname: '',
            avatar: '',
            gender: '',
            voiceTag: '',
            interestTags: '',
            level: undefined,
            vipStatus: '',
            vipExpireTime: new Date(),
            balance: 0,
            diamond: undefined,
            registerTime: new Date(),
            lastLoginTime: new Date(),
            lastLoginIp: '',
            status: '',
            freezeReason: '',
            freezeEndTime: new Date(),
            isAuth: '',
        })
// 验证规则
const rule = reactive({
               username : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               password : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               nickname : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               avatar : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               gender : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               level : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               vipStatus : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               balance : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               diamond : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               registerTime : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               lastLoginTime : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               lastLoginIp : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               status : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               isAuth : [{
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
      const res = await findUser({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    genderOptions.value = await getDictFunc('gender')
    voiceOptions.value = await getDictFunc('voice')
    interesttagsOptions.value = await getDictFunc('interesttags')
    userstatusOptions.value = await getDictFunc('userstatus')
    vipstatusOptions.value = await getDictFunc('vipstatus')
    is_authOptions.value = await getDictFunc('is_auth')
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
               res = await createUser(formData.value)
               break
             case 'update':
               res = await updateUser(formData.value)
               break
             default:
               res = await createUser(formData.value)
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
