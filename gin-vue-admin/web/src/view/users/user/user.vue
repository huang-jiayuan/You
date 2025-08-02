
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAtRange">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>

      <el-date-picker
            v-model="searchInfo.createdAtRange"
            class="w-[380px]"
            type="datetimerange"
            range-separator="至"
            start-placeholder="开始时间"
            end-placeholder="结束时间"
          />
       </el-form-item>
      
            <el-form-item label="登录账号(手机号/第三方账号)" prop="username">
  <el-input v-model="searchInfo.username" placeholder="搜索条件" />
</el-form-item>
            
            <el-form-item label="登录密码(SHA-256加密)" prop="password">
  <el-input v-model="searchInfo.password" placeholder="搜索条件" />
</el-form-item>
            
            <el-form-item label="用户昵称" prop="nickname">
  <el-input v-model="searchInfo.nickname" placeholder="搜索条件" />
</el-form-item>
            
            <el-form-item label="头像URL" prop="avatar">
  <el-input v-model="searchInfo.avatar" placeholder="搜索条件" />
</el-form-item>
            
            <el-form-item label="性别(0-保密,1-男,2-女)" prop="gender">
  <el-select v-model="searchInfo.gender" clearable filterable placeholder="请选择" @clear="()=>{searchInfo.gender=undefined}">
    <el-option v-for="(item,key) in genderOptions" :key="key" :label="item.label" :value="item.value" />
  </el-select>
</el-form-item>
            
            <el-form-item label="声音标签(逗号分隔)" prop="voiceTag">
  <el-select v-model="searchInfo.voiceTag" clearable filterable placeholder="请选择" @clear="()=>{searchInfo.voiceTag=undefined}">
    <el-option v-for="(item,key) in voiceOptions" :key="key" :label="item.label" :value="item.value" />
  </el-select>
</el-form-item>
            
            <el-form-item label="兴趣标签(逗号分隔)" prop="interestTags">
  <el-select v-model="searchInfo.interestTags" clearable filterable placeholder="请选择" @clear="()=>{searchInfo.interestTags=undefined}">
    <el-option v-for="(item,key) in interesttagsOptions" :key="key" :label="item.label" :value="item.value" />
  </el-select>
</el-form-item>
            
            <el-form-item label="用户等级" prop="level">
  <el-input v-model.number="searchInfo.level" placeholder="搜索条件" />
</el-form-item>
            
            <el-form-item label="会员状态(0-非会员,1-月会员,2-年会员)" prop="vipStatus">
  <el-select v-model="searchInfo.vipStatus" clearable filterable placeholder="请选择" @clear="()=>{searchInfo.vipStatus=undefined}">
    <el-option v-for="(item,key) in vipstatusOptions" :key="key" :label="item.label" :value="item.value" />
  </el-select>
</el-form-item>
            
            <el-form-item label="会员过期时间" prop="vipExpireTime">
<el-date-picker v-model="searchInfo.vipExpireTime" type="datetime" placeholder="搜索条件"></el-date-picker></el-form-item>
            
            <el-form-item label="虚拟币余额" prop="balance">
  <el-input v-model.number="searchInfo.balance" placeholder="搜索条件" />
</el-form-item>
            
            <el-form-item label="钻石数量" prop="diamond">
  <el-input v-model.number="searchInfo.diamond" placeholder="搜索条件" />
</el-form-item>
            
            <el-form-item label="最后登录IP" prop="lastLoginIp">
  <el-input v-model="searchInfo.lastLoginIp" placeholder="搜索条件" />
</el-form-item>
            
            <el-form-item label="账号状态(0-正常,1-封禁,2-冻结,3-禁言)" prop="status">
  <el-select v-model="searchInfo.status" clearable filterable placeholder="请选择" @clear="()=>{searchInfo.status=undefined}">
    <el-option v-for="(item,key) in userstatusOptions" :key="key" :label="item.label" :value="item.value" />
  </el-select>
</el-form-item>
            
            <el-form-item label="冻结/封禁原因" prop="freezeReason">
  <el-input v-model="searchInfo.freezeReason" placeholder="搜索条件" />
</el-form-item>
            
            <el-form-item label="冻结/封禁结束时间" prop="freezeEndTime">
<el-date-picker v-model="searchInfo.freezeEndTime" type="datetime" placeholder="搜索条件"></el-date-picker></el-form-item>
            
            <el-form-item label="实名认证状态（0：未认证1：已认证2）" prop="isAuth">
  <el-select v-model="searchInfo.isAuth" clearable filterable placeholder="请选择" @clear="()=>{searchInfo.isAuth=undefined}">
    <el-option v-for="(item,key) in is_authOptions" :key="key" :label="item.label" :value="item.value" />
  </el-select>
</el-form-item>
            

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button v-auth="btnAuth.add" type="primary" icon="plus" @click="openDialog()">新增</el-button>
            <el-button v-auth="btnAuth.batchDelete" icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
            <ExportTemplate v-auth="btnAuth.exportTemplate" template-id="users_User" />
            <ExportExcel v-auth="btnAuth.exportExcel" template-id="users_User" filterDeleted/>
            <ImportExcel v-auth="btnAuth.importExcel" template-id="users_User" @on-success="getTableData" />
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column sortable align="left" label="日期" prop="CreatedAt"width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
            <el-table-column align="left" label="登录账号(手机号/第三方账号)" prop="username" width="120" />

            <el-table-column align="left" label="登录密码(SHA-256加密)" prop="password" width="120" />

            <el-table-column align="left" label="用户昵称" prop="nickname" width="120" />

            <el-table-column align="left" label="头像URL" prop="avatar" width="120" />

            <el-table-column align="left" label="性别(0-保密,1-男,2-女)" prop="gender" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.gender,genderOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="声音标签(逗号分隔)" prop="voiceTag" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.voiceTag,voiceOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="兴趣标签(逗号分隔)" prop="interestTags" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.interestTags,interesttagsOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="用户等级" prop="level" width="120" />

            <el-table-column align="left" label="会员状态(0-非会员,1-月会员,2-年会员)" prop="vipStatus" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.vipStatus,vipstatusOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="会员过期时间" prop="vipExpireTime" width="180">
   <template #default="scope">{{ formatDate(scope.row.vipExpireTime) }}</template>
</el-table-column>
            <el-table-column align="left" label="虚拟币余额" prop="balance" width="120" />

            <el-table-column align="left" label="钻石数量" prop="diamond" width="120" />

            <el-table-column align="left" label="注册时间" prop="registerTime" width="180">
   <template #default="scope">{{ formatDate(scope.row.registerTime) }}</template>
</el-table-column>
            <el-table-column align="left" label="最后登录时间" prop="lastLoginTime" width="180">
   <template #default="scope">{{ formatDate(scope.row.lastLoginTime) }}</template>
</el-table-column>
            <el-table-column align="left" label="最后登录IP" prop="lastLoginIp" width="120" />

            <el-table-column align="left" label="账号状态(0-正常,1-封禁,2-冻结,3-禁言)" prop="status" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.status,userstatusOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="冻结/封禁原因" prop="freezeReason" width="120" />

            <el-table-column align="left" label="冻结/封禁结束时间" prop="freezeEndTime" width="180">
   <template #default="scope">{{ formatDate(scope.row.freezeEndTime) }}</template>
</el-table-column>
            <el-table-column align="left" label="实名认证状态（0：未认证1：已认证2）" prop="isAuth" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.isAuth,is_authOptions) }}
    </template>
</el-table-column>
        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
            <template #default="scope">
            <el-button v-auth="btnAuth.info" type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button v-auth="btnAuth.edit" type="primary" link icon="edit" class="table-button" @click="updateUserFunc(scope.row)">编辑</el-button>
            <el-button  v-auth="btnAuth.delete" type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'新增':'编辑'}}</span>
                <div>
                  <el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
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
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="登录账号(手机号/第三方账号)">
    {{ detailFrom.username }}
</el-descriptions-item>
                    <el-descriptions-item label="登录密码(SHA-256加密)">
    {{ detailFrom.password }}
</el-descriptions-item>
                    <el-descriptions-item label="用户昵称">
    {{ detailFrom.nickname }}
</el-descriptions-item>
                    <el-descriptions-item label="头像URL">
    {{ detailFrom.avatar }}
</el-descriptions-item>
                    <el-descriptions-item label="性别(0-保密,1-男,2-女)">
    {{ detailFrom.gender }}
</el-descriptions-item>
                    <el-descriptions-item label="声音标签(逗号分隔)">
    {{ detailFrom.voiceTag }}
</el-descriptions-item>
                    <el-descriptions-item label="兴趣标签(逗号分隔)">
    {{ detailFrom.interestTags }}
</el-descriptions-item>
                    <el-descriptions-item label="用户等级">
    {{ detailFrom.level }}
</el-descriptions-item>
                    <el-descriptions-item label="会员状态(0-非会员,1-月会员,2-年会员)">
    {{ detailFrom.vipStatus }}
</el-descriptions-item>
                    <el-descriptions-item label="会员过期时间">
    {{ detailFrom.vipExpireTime }}
</el-descriptions-item>
                    <el-descriptions-item label="虚拟币余额">
    {{ detailFrom.balance }}
</el-descriptions-item>
                    <el-descriptions-item label="钻石数量">
    {{ detailFrom.diamond }}
</el-descriptions-item>
                    <el-descriptions-item label="注册时间">
    {{ detailFrom.registerTime }}
</el-descriptions-item>
                    <el-descriptions-item label="最后登录时间">
    {{ detailFrom.lastLoginTime }}
</el-descriptions-item>
                    <el-descriptions-item label="最后登录IP">
    {{ detailFrom.lastLoginIp }}
</el-descriptions-item>
                    <el-descriptions-item label="账号状态(0-正常,1-封禁,2-冻结,3-禁言)">
    {{ detailFrom.status }}
</el-descriptions-item>
                    <el-descriptions-item label="冻结/封禁原因">
    {{ detailFrom.freezeReason }}
</el-descriptions-item>
                    <el-descriptions-item label="冻结/封禁结束时间">
    {{ detailFrom.freezeEndTime }}
</el-descriptions-item>
                    <el-descriptions-item label="实名认证状态（0：未认证1：已认证2）">
    {{ detailFrom.isAuth }}
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createUser,
  deleteUser,
  deleteUserByIds,
  updateUser,
  findUser,
  getUserList
} from '@/api/users/user'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
// 引入按钮权限标识
import { useBtnAuth } from '@/utils/btnAuth'
import { useAppStore } from "@/pinia"

// 导出组件
import ExportExcel from '@/components/exportExcel/exportExcel.vue'
// 导入组件
import ImportExcel from '@/components/exportExcel/importExcel.vue'
// 导出模板组件
import ExportTemplate from '@/components/exportExcel/exportTemplate.vue'


defineOptions({
    name: 'User'
})
// 按钮权限实例化
    const btnAuth = useBtnAuth()

// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
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
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               password : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               nickname : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               avatar : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               gender : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               level : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               vipStatus : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               balance : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               diamond : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               registerTime : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               lastLoginTime : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               lastLoginIp : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               status : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               isAuth : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getUserList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    genderOptions.value = await getDictFunc('gender')
    voiceOptions.value = await getDictFunc('voice')
    interesttagsOptions.value = await getDictFunc('interesttags')
    userstatusOptions.value = await getDictFunc('userstatus')
    vipstatusOptions.value = await getDictFunc('vipstatus')
    is_authOptions.value = await getDictFunc('is_auth')
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteUserFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
      const res = await deleteUserByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateUserFunc = async(row) => {
    const res = await findUser({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteUserFunc = async (row) => {
    const res = await deleteUser({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
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
        }
}
// 弹窗确定
const enterDialog = async () => {
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
                closeDialog()
                getTableData()
              }
      })
}

const detailFrom = ref({})

// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findUser({ ID: row.ID })
  if (res.code === 0) {
    detailFrom.value = res.data
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailFrom.value = {}
}


</script>

<style>

</style>
