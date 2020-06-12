<p>Packages:</p>
<ul>
<li>
<a href="#onecloud.yunion.io%2fv1">onecloud.yunion.io/v1</a>
</li>
</ul>
<h2 id="onecloud.yunion.io/v1">onecloud.yunion.io/v1</h2>
<p>
</p>
Resource Types:
<ul></ul>
<h3 id="onecloud.yunion.io/v1.AnsiblePlaybook">AnsiblePlaybook
</h3>
<p>
<p>AnsiblePlaybook is the Schema for the ansibleplaybooks API</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metadata</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
<em>(Required)</em>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code></br>
<em>
<a href="#onecloud.yunion.io/v1.AnsiblePlaybookSpec">
AnsiblePlaybookSpec
</a>
</em>
</td>
<td>
<em>(Required)</em>
<br/>
<br/>
<table>
<tr>
<td>
<code>playbookTemplateRef</code></br>
<em>
<a href="#onecloud.yunion.io/v1.LocalObjectReference">
LocalObjectReference
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>PlaybookTemplateRef specifies the AnsiblePlaybookTemplate.</p>
</td>
</tr>
<tr>
<td>
<code>playbookTemplate</code></br>
<em>
<a href="#onecloud.yunion.io/v1.AnsiblePlaybookTemplateSpec">
AnsiblePlaybookTemplateSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>PlaybookTemplate describe the ansible playbook</p>
</td>
</tr>
<tr>
<td>
<code>inventory</code></br>
<em>
<a href="#onecloud.yunion.io/v1.AnsiblePlaybookHost">
[]AnsiblePlaybookHost
</a>
</em>
</td>
<td>
<em>(Required)</em>
<p>VirtualMachines specifies the inventory of ansible playbook.</p>
</td>
</tr>
<tr>
<td>
<code>vars</code></br>
<em>
<a href="#onecloud.yunion.io/v1.IntOrStringStore">
map[string]../api/v1.IntOrStringStore
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Vars describe the public value about Vars in AnsiblePlaybookTemplate.</p>
</td>
</tr>
<tr>
<td>
<code>maxRetryTimes</code></br>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Nil or Non-positive number means unlimited.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code></br>
<em>
<a href="#onecloud.yunion.io/v1.AnsiblePlaybookStatus">
AnsiblePlaybookStatus
</a>
</em>
</td>
<td>
<em>(Required)</em>
</td>
</tr>
</tbody>
</table>
<h3 id="onecloud.yunion.io/v1.AnsiblePlaybookHost">AnsiblePlaybookHost
</h3>
<p>
(<em>Appears on:</em>
<a href="#onecloud.yunion.io/v1.AnsiblePlaybookSpec">AnsiblePlaybookSpec</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>virtualMachine</code></br>
<em>
<a href="#onecloud.yunion.io/v1.ObjectReference">
ObjectReference
</a>
</em>
</td>
<td>
<em>(Required)</em>
</td>
</tr>
<tr>
<td>
<code>vars</code></br>
<em>
<a href="#onecloud.yunion.io/v1.IntOrStringStore">
map[string]../api/v1.IntOrStringStore
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Vars describes the unique values ​​of the VirtualMachine
corresponding to the variables in the AnsiblePlaybookTemplate.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="onecloud.yunion.io/v1.AnsiblePlaybookInfo">AnsiblePlaybookInfo
</h3>
<p>
(<em>Appears on:</em>
<a href="#onecloud.yunion.io/v1.AnsiblePlaybookStatus">AnsiblePlaybookStatus</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ExternalInfoBase</code></br>
<em>
<a href="#onecloud.yunion.io/v1.ExternalInfoBase">
ExternalInfoBase
</a>
</em>
</td>
<td>
<p>
(Members of <code>ExternalInfoBase</code> are embedded into this type.)
</p>
<em>(Required)</em>
</td>
</tr>
<tr>
<td>
<code>output</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
</tbody>
</table>
<h3 id="onecloud.yunion.io/v1.AnsiblePlaybookSpec">AnsiblePlaybookSpec
</h3>
<p>
(<em>Appears on:</em>
<a href="#onecloud.yunion.io/v1.AnsiblePlaybook">AnsiblePlaybook</a>)
</p>
<p>
<p>AnsiblePlaybookSpec defines the desired state of AnsiblePlaybook</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>playbookTemplateRef</code></br>
<em>
<a href="#onecloud.yunion.io/v1.LocalObjectReference">
LocalObjectReference
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>PlaybookTemplateRef specifies the AnsiblePlaybookTemplate.</p>
</td>
</tr>
<tr>
<td>
<code>playbookTemplate</code></br>
<em>
<a href="#onecloud.yunion.io/v1.AnsiblePlaybookTemplateSpec">
AnsiblePlaybookTemplateSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>PlaybookTemplate describe the ansible playbook</p>
</td>
</tr>
<tr>
<td>
<code>inventory</code></br>
<em>
<a href="#onecloud.yunion.io/v1.AnsiblePlaybookHost">
[]AnsiblePlaybookHost
</a>
</em>
</td>
<td>
<em>(Required)</em>
<p>VirtualMachines specifies the inventory of ansible playbook.</p>
</td>
</tr>
<tr>
<td>
<code>vars</code></br>
<em>
<a href="#onecloud.yunion.io/v1.IntOrStringStore">
map[string]../api/v1.IntOrStringStore
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Vars describe the public value about Vars in AnsiblePlaybookTemplate.</p>
</td>
</tr>
<tr>
<td>
<code>maxRetryTimes</code></br>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Nil or Non-positive number means unlimited.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="onecloud.yunion.io/v1.AnsiblePlaybookStatus">AnsiblePlaybookStatus
</h3>
<p>
(<em>Appears on:</em>
<a href="#onecloud.yunion.io/v1.AnsiblePlaybook">AnsiblePlaybook</a>)
</p>
<p>
<p>AnsiblePlaybookStatus defines the observed state of AnsiblePlaybook</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>phase</code></br>
<em>
<a href="#onecloud.yunion.io/v1.ResourcePhase">
ResourcePhase
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Important: Run &ldquo;make&rdquo; to regenerate code after modifying this file</p>
</td>
</tr>
<tr>
<td>
<code>reason</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>A human readable message indicating details about why vm is in this phase.</p>
</td>
</tr>
<tr>
<td>
<code>externalInfo</code></br>
<em>
<a href="#onecloud.yunion.io/v1.AnsiblePlaybookInfo">
AnsiblePlaybookInfo
</a>
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
<tr>
<td>
<code>tryTimes</code></br>
<em>
int32
</em>
</td>
<td>
<em>(Required)</em>
</td>
</tr>
</tbody>
</table>
<h3 id="onecloud.yunion.io/v1.AnsiblePlaybookTemplate">AnsiblePlaybookTemplate
</h3>
<p>
<p>AnsiblePlaybookTemplate is the Schema for the ansibleplaybooktemplates API</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metadata</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
<em>(Required)</em>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code></br>
<em>
<a href="#onecloud.yunion.io/v1.AnsiblePlaybookTemplateSpec">
AnsiblePlaybookTemplateSpec
</a>
</em>
</td>
<td>
<em>(Required)</em>
<br/>
<br/>
<table>
<tr>
<td>
<code>playbook</code></br>
<em>
string
</em>
</td>
<td>
<em>(Required)</em>
<p>Playbook describe the main content of absible playbook which should be in yaml format.</p>
</td>
</tr>
<tr>
<td>
<code>requirements</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Requirements describe the source of roles dependent on Playbook</p>
</td>
</tr>
<tr>
<td>
<code>files</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Files describe the associated file tree and file content which should be in json format.</p>
</td>
</tr>
<tr>
<td>
<code>vars</code></br>
<em>
<a href="#onecloud.yunion.io/v1.AnsiblePlaybookTemplateVar">
[]AnsiblePlaybookTemplateVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Vars describe the vars to apply this ansible playbook.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code></br>
<em>
<a href="#onecloud.yunion.io/v1.AnsiblePlaybookTemplateStatus">
AnsiblePlaybookTemplateStatus
</a>
</em>
</td>
<td>
<em>(Required)</em>
</td>
</tr>
</tbody>
</table>
<h3 id="onecloud.yunion.io/v1.AnsiblePlaybookTemplateSpec">AnsiblePlaybookTemplateSpec
</h3>
<p>
(<em>Appears on:</em>
<a href="#onecloud.yunion.io/v1.AnsiblePlaybookSpec">AnsiblePlaybookSpec</a>, 
<a href="#onecloud.yunion.io/v1.AnsiblePlaybookTemplate">AnsiblePlaybookTemplate</a>)
</p>
<p>
<p>AnsiblePlaybookTemplateSpec defines the desired state of AnsiblePlaybookTemplate.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>playbook</code></br>
<em>
string
</em>
</td>
<td>
<em>(Required)</em>
<p>Playbook describe the main content of absible playbook which should be in yaml format.</p>
</td>
</tr>
<tr>
<td>
<code>requirements</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Requirements describe the source of roles dependent on Playbook</p>
</td>
</tr>
<tr>
<td>
<code>files</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Files describe the associated file tree and file content which should be in json format.</p>
</td>
</tr>
<tr>
<td>
<code>vars</code></br>
<em>
<a href="#onecloud.yunion.io/v1.AnsiblePlaybookTemplateVar">
[]AnsiblePlaybookTemplateVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Vars describe the vars to apply this ansible playbook.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="onecloud.yunion.io/v1.AnsiblePlaybookTemplateStatus">AnsiblePlaybookTemplateStatus
</h3>
<p>
(<em>Appears on:</em>
<a href="#onecloud.yunion.io/v1.AnsiblePlaybookTemplate">AnsiblePlaybookTemplate</a>)
</p>
<p>
<p>AnsiblePlaybookTemplateStatus defines the observed state of AnsiblePlaybookTemplate</p>
</p>
<h3 id="onecloud.yunion.io/v1.AnsiblePlaybookTemplateVar">AnsiblePlaybookTemplateVar
</h3>
<p>
(<em>Appears on:</em>
<a href="#onecloud.yunion.io/v1.AnsiblePlaybookTemplateSpec">AnsiblePlaybookTemplateSpec</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code></br>
<em>
string
</em>
</td>
<td>
<em>(Required)</em>
</td>
</tr>
<tr>
<td>
<code>required</code></br>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Required indicates whether this variable is required.</p>
</td>
</tr>
<tr>
<td>
<code>default</code></br>
<em>
<a href="#onecloud.yunion.io/v1.IntOrString">
IntOrString
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Default describe the default value of this variable.
If it is empty, Required should be true.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="onecloud.yunion.io/v1.DiskDriver">DiskDriver
(<code>string</code> alias)</p></h3>
<p>
(<em>Appears on:</em>
<a href="#onecloud.yunion.io/v1.VMDiskSpec">VMDiskSpec</a>)
</p>
<p>
<p>Driver represents the drive method of the disk on the virtual machine.</p>
</p>
<h3 id="onecloud.yunion.io/v1.DisplayIdenSpec">DisplayIdenSpec
</h3>
<p>
(<em>Appears on:</em>
<a href="#onecloud.yunion.io/v1.VirtualMachineSpec">VirtualMachineSpec</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Name</p>
</td>
</tr>
<tr>
<td>
<code>nameCheck</code></br>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>NameCheck indicates whether to strictly check Name</p>
</td>
</tr>
<tr>
<td>
<code>description</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
</tbody>
</table>
<h3 id="onecloud.yunion.io/v1.ExternalInfoBase">ExternalInfoBase
</h3>
<p>
(<em>Appears on:</em>
<a href="#onecloud.yunion.io/v1.AnsiblePlaybookInfo">AnsiblePlaybookInfo</a>, 
<a href="#onecloud.yunion.io/v1.VMInfo">VMInfo</a>)
</p>
<p>
<p>ExternalInfoBase describe the corresponding resource&rsquo;s info in external system</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>id</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
<tr>
<td>
<code>status</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
<tr>
<td>
<code>action</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Action indicate the latest action for external vm.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="onecloud.yunion.io/v1.IResource">IResource
</h3>
<p>
</p>
<h3 id="onecloud.yunion.io/v1.IStore">IStore
</h3>
<p>
</p>
<h3 id="onecloud.yunion.io/v1.IValue">IValue
</h3>
<p>
</p>
<h3 id="onecloud.yunion.io/v1.IntOrString">IntOrString
</h3>
<p>
(<em>Appears on:</em>
<a href="#onecloud.yunion.io/v1.AnsiblePlaybookTemplateVar">AnsiblePlaybookTemplateVar</a>, 
<a href="#onecloud.yunion.io/v1.IntOrStringStore">IntOrStringStore</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>IntOrString</code></br>
<em>
<a href="https://godoc.org/k8s.io/apimachinery/pkg/util/intstr#IntOrString">
k8s.io/apimachinery/pkg/util/intstr.IntOrString
</a>
</em>
</td>
<td>
<p>
(Members of <code>IntOrString</code> are embedded into this type.)
</p>
<em>(Required)</em>
</td>
</tr>
</tbody>
</table>
<h3 id="onecloud.yunion.io/v1.IntOrStringStore">IntOrStringStore
</h3>
<p>
(<em>Appears on:</em>
<a href="#onecloud.yunion.io/v1.AnsiblePlaybookHost">AnsiblePlaybookHost</a>, 
<a href="#onecloud.yunion.io/v1.AnsiblePlaybookSpec">AnsiblePlaybookSpec</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>value</code></br>
<em>
<a href="#onecloud.yunion.io/v1.IntOrString">
IntOrString
</a>
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
<tr>
<td>
<code>reference</code></br>
<em>
<a href="#onecloud.yunion.io/v1.ObjectFieldReference">
ObjectFieldReference
</a>
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
</tbody>
</table>
<h3 id="onecloud.yunion.io/v1.LocalObjectReference">LocalObjectReference
</h3>
<p>
(<em>Appears on:</em>
<a href="#onecloud.yunion.io/v1.AnsiblePlaybookSpec">AnsiblePlaybookSpec</a>)
</p>
<p>
<p>LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code></br>
<em>
string
</em>
</td>
<td>
<em>(Required)</em>
</td>
</tr>
</tbody>
</table>
<h3 id="onecloud.yunion.io/v1.ObjectFieldReference">ObjectFieldReference
</h3>
<p>
(<em>Appears on:</em>
<a href="#onecloud.yunion.io/v1.IntOrStringStore">IntOrStringStore</a>, 
<a href="#onecloud.yunion.io/v1.StringStore">StringStore</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>group</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
<tr>
<td>
<code>version</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
<tr>
<td>
<code>kind</code></br>
<em>
string
</em>
</td>
<td>
<em>(Required)</em>
</td>
</tr>
<tr>
<td>
<code>namespace</code></br>
<em>
string
</em>
</td>
<td>
<em>(Required)</em>
</td>
</tr>
<tr>
<td>
<code>name</code></br>
<em>
string
</em>
</td>
<td>
<em>(Required)</em>
</td>
</tr>
<tr>
<td>
<code>fieldPath</code></br>
<em>
string
</em>
</td>
<td>
<em>(Required)</em>
</td>
</tr>
</tbody>
</table>
<h3 id="onecloud.yunion.io/v1.ObjectReference">ObjectReference
</h3>
<p>
(<em>Appears on:</em>
<a href="#onecloud.yunion.io/v1.AnsiblePlaybookHost">AnsiblePlaybookHost</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>kind</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
<tr>
<td>
<code>namespace</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
<tr>
<td>
<code>name</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
<tr>
<td>
<code>uid</code></br>
<em>
<a href="https://godoc.org/k8s.io/apimachinery/pkg/types#UID">
k8s.io/apimachinery/pkg/types.UID
</a>
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
<tr>
<td>
<code>apiVersion</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
<tr>
<td>
<code>resourceVersion</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
</tbody>
</table>
<h3 id="onecloud.yunion.io/v1.RecreatePolicy">RecreatePolicy
</h3>
<p>
(<em>Appears on:</em>
<a href="#onecloud.yunion.io/v1.VirtualMachineSpec">VirtualMachineSpec</a>)
</p>
<p>
<p>RecreatePolicy describe that when the virtual machine is abnormal, how to deal with it,
specifically determine whether to delete and recreate.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>matchStatus</code></br>
<em>
[]string
</em>
</td>
<td>
<em>(Required)</em>
</td>
</tr>
<tr>
<td>
<code>never</code></br>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
<tr>
<td>
<code>allways</code></br>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
<tr>
<td>
<code>maxTimes</code></br>
<em>
int32
</em>
</td>
<td>
<em>(Required)</em>
</td>
</tr>
</tbody>
</table>
<h3 id="onecloud.yunion.io/v1.ResourcePhase">ResourcePhase
(<code>string</code> alias)</p></h3>
<p>
(<em>Appears on:</em>
<a href="#onecloud.yunion.io/v1.AnsiblePlaybookStatus">AnsiblePlaybookStatus</a>, 
<a href="#onecloud.yunion.io/v1.VirtualMachineStatus">VirtualMachineStatus</a>)
</p>
<p>
<p>ResourcePhase is a label for the condition of a resource at the current time</p>
</p>
<h3 id="onecloud.yunion.io/v1.StorageMedium">StorageMedium
(<code>string</code> alias)</p></h3>
<p>
(<em>Appears on:</em>
<a href="#onecloud.yunion.io/v1.VMStorageSpec">VMStorageSpec</a>)
</p>
<p>
<p>StorageMedium represents storage media type</p>
</p>
<h3 id="onecloud.yunion.io/v1.String">String
(<code>string</code> alias)</p></h3>
<p>
</p>
<h3 id="onecloud.yunion.io/v1.StringStore">StringStore
</h3>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>value</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
<tr>
<td>
<code>reference</code></br>
<em>
<a href="#onecloud.yunion.io/v1.ObjectFieldReference">
ObjectFieldReference
</a>
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
</tbody>
</table>
<h3 id="onecloud.yunion.io/v1.VMDiskSpec">VMDiskSpec
</h3>
<p>
(<em>Appears on:</em>
<a href="#onecloud.yunion.io/v1.VirtualMachineConfig">VirtualMachineConfig</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>image</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>The disk will be created from the image represented by ImageId.</p>
</td>
</tr>
<tr>
<td>
<code>sizeGB</code></br>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>SizeGB represents the size(unit: GB) of disk.</p>
</td>
</tr>
<tr>
<td>
<code>driver</code></br>
<em>
<a href="#onecloud.yunion.io/v1.DiskDriver">
DiskDriver
</a>
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
<tr>
<td>
<code>storageConfig</code></br>
<em>
<a href="#onecloud.yunion.io/v1.VMStorageSpec">
VMStorageSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
</tbody>
</table>
<h3 id="onecloud.yunion.io/v1.VMInfo">VMInfo
</h3>
<p>
(<em>Appears on:</em>
<a href="#onecloud.yunion.io/v1.VirtualMachineStatus">VirtualMachineStatus</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ExternalInfoBase</code></br>
<em>
<a href="#onecloud.yunion.io/v1.ExternalInfoBase">
ExternalInfoBase
</a>
</em>
</td>
<td>
<p>
(Members of <code>ExternalInfoBase</code> are embedded into this type.)
</p>
<em>(Required)</em>
</td>
</tr>
<tr>
<td>
<code>eip</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
<tr>
<td>
<code>ips</code></br>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
</tbody>
</table>
<h3 id="onecloud.yunion.io/v1.VMNetworkSpec">VMNetworkSpec
</h3>
<p>
(<em>Appears on:</em>
<a href="#onecloud.yunion.io/v1.VirtualMachineConfig">VirtualMachineConfig</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>network</code></br>
<em>
string
</em>
</td>
<td>
<em>(Required)</em>
</td>
</tr>
<tr>
<td>
<code>address</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
</tbody>
</table>
<h3 id="onecloud.yunion.io/v1.VMNewEipSpec">VMNewEipSpec
</h3>
<p>
(<em>Appears on:</em>
<a href="#onecloud.yunion.io/v1.VirtualMachineSpec">VirtualMachineSpec</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>bw</code></br>
<em>
int64
</em>
</td>
<td>
<em>(Required)</em>
<p>Bw indicates the bandwidth of the Elastic Public IP.</p>
</td>
</tr>
<tr>
<td>
<code>chargeType</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>The charge type of Elastic Public IP</p>
</td>
</tr>
</tbody>
</table>
<h3 id="onecloud.yunion.io/v1.VMPasswordSpec">VMPasswordSpec
</h3>
<p>
(<em>Appears on:</em>
<a href="#onecloud.yunion.io/v1.VirtualMachineSpec">VirtualMachineSpec</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>keyPairId</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
<tr>
<td>
<code>password</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
<tr>
<td>
<code>resetPassword</code></br>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
</tbody>
</table>
<h3 id="onecloud.yunion.io/v1.VMPreferSpec">VMPreferSpec
</h3>
<p>
(<em>Appears on:</em>
<a href="#onecloud.yunion.io/v1.VirtualMachineConfig">VirtualMachineConfig</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>preferManager</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>PreferManager is the parameter passed to the scheduler which makes
the virtual machine created in the preferred cloud provider.</p>
</td>
</tr>
<tr>
<td>
<code>preferRegion</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>PreferRegion is the parameter passed to the scheduler which makes
the virtual machine created in the preferred cloud region.</p>
</td>
</tr>
<tr>
<td>
<code>preferZone</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>PreferZone is the parameter passed to the scheduler which makes
the virtual machine created in the preferred cloud zone.</p>
</td>
</tr>
<tr>
<td>
<code>preferWire</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>PreferWire is the parameter passed to the scheduler which makes
the machine created in the preferred wire.</p>
</td>
</tr>
<tr>
<td>
<code>preferHost</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>PreferHost is the parameter passed to the scheduler which makes
the machine created in the preferred host.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="onecloud.yunion.io/v1.VMProjectSpec">VMProjectSpec
</h3>
<p>
(<em>Appears on:</em>
<a href="#onecloud.yunion.io/v1.VirtualMachineSpec">VirtualMachineSpec</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>project</code></br>
<em>
string
</em>
</td>
<td>
<em>(Required)</em>
</td>
</tr>
<tr>
<td>
<code>projectDomain</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
</tbody>
</table>
<h3 id="onecloud.yunion.io/v1.VMStorageSpec">VMStorageSpec
</h3>
<p>
(<em>Appears on:</em>
<a href="#onecloud.yunion.io/v1.VMDiskSpec">VMDiskSpec</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>storage</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Storage represents specific storage</p>
</td>
</tr>
<tr>
<td>
<code>backend</code></br>
<em>
string
</em>
</td>
<td>
<em>(Required)</em>
<p>Backend represents backend of storage</p>
</td>
</tr>
<tr>
<td>
<code>medium</code></br>
<em>
<a href="#onecloud.yunion.io/v1.StorageMedium">
StorageMedium
</a>
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
</tbody>
</table>
<h3 id="onecloud.yunion.io/v1.VirtualMachine">VirtualMachine
</h3>
<p>
<p>VirtualMachine is the Schema for the virtualmachines API</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metadata</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
<em>(Required)</em>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code></br>
<em>
<a href="#onecloud.yunion.io/v1.VirtualMachineSpec">
VirtualMachineSpec
</a>
</em>
</td>
<td>
<em>(Required)</em>
<br/>
<br/>
<table>
<tr>
<td>
<code>DisplayIdenSpec</code></br>
<em>
<a href="#onecloud.yunion.io/v1.DisplayIdenSpec">
DisplayIdenSpec
</a>
</em>
</td>
<td>
<p>
(Members of <code>DisplayIdenSpec</code> are embedded into this type.)
</p>
<em>(Required)</em>
</td>
</tr>
<tr>
<td>
<code>vmConfig</code></br>
<em>
<a href="#onecloud.yunion.io/v1.VirtualMachineConfig">
VirtualMachineConfig
</a>
</em>
</td>
<td>
<em>(Required)</em>
</td>
</tr>
<tr>
<td>
<code>projectConfig</code></br>
<em>
<a href="#onecloud.yunion.io/v1.VMProjectSpec">
VMProjectSpec
</a>
</em>
</td>
<td>
<em>(Required)</em>
</td>
</tr>
<tr>
<td>
<code>VMPasswordSpec</code></br>
<em>
<a href="#onecloud.yunion.io/v1.VMPasswordSpec">
VMPasswordSpec
</a>
</em>
</td>
<td>
<p>
(Members of <code>VMPasswordSpec</code> are embedded into this type.)
</p>
<em>(Required)</em>
</td>
</tr>
<tr>
<td>
<code>secgroups</code></br>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
<tr>
<td>
<code>newEip</code></br>
<em>
<a href="#onecloud.yunion.io/v1.VMNewEipSpec">
VMNewEipSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>NewEip indicates that create a new EIP and bind it with VM.
It conflicts with Eip.</p>
</td>
</tr>
<tr>
<td>
<code>eip</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Eip indicates that bind the existing EIP with VM.
It conflicts with NewEip.</p>
</td>
</tr>
<tr>
<td>
<code>billDuration</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>BillDuration describes the duration of the annual and monthly billing type.
That length of BillDuration represents the billing type is &lsquo;prepaid&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>autoRenew</code></br>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>AutoRenew indicates whether to automatically renewal.</p>
</td>
</tr>
<tr>
<td>
<code>recreatePolicy</code></br>
<em>
<a href="#onecloud.yunion.io/v1.RecreatePolicy">
RecreatePolicy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code></br>
<em>
<a href="#onecloud.yunion.io/v1.VirtualMachineStatus">
VirtualMachineStatus
</a>
</em>
</td>
<td>
<em>(Required)</em>
</td>
</tr>
</tbody>
</table>
<h3 id="onecloud.yunion.io/v1.VirtualMachineConfig">VirtualMachineConfig
</h3>
<p>
(<em>Appears on:</em>
<a href="#onecloud.yunion.io/v1.VirtualMachineSpec">VirtualMachineSpec</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>VMPreferSpec</code></br>
<em>
<a href="#onecloud.yunion.io/v1.VMPreferSpec">
VMPreferSpec
</a>
</em>
</td>
<td>
<p>
(Members of <code>VMPreferSpec</code> are embedded into this type.)
</p>
<em>(Required)</em>
</td>
</tr>
<tr>
<td>
<code>hypervisor</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
<tr>
<td>
<code>vcpuCount</code></br>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>VcpuCount represents the number of CPUs of the virtual machine.
It conflicts with InstanceType and it is It is required if InstanceType is not specified.</p>
</td>
</tr>
<tr>
<td>
<code>vmemSizeGB</code></br>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>VmemSizeGB represents the size of memory of the virtual machine.
It conflicts with InstanceType and it is It is required if InstanceType is not specified.</p>
</td>
</tr>
<tr>
<td>
<code>instanceType</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>InstanceType describes the specifications of the virtual machine,
which are predefined by the cloud provider.
It conflicts with VcpuCount and VmemSizeGB.</p>
</td>
</tr>
<tr>
<td>
<code>rootDisk</code></br>
<em>
<a href="#onecloud.yunion.io/v1.VMDiskSpec">
VMDiskSpec
</a>
</em>
</td>
<td>
<em>(Required)</em>
<p>RootDisk describes the configuration of the system disk</p>
</td>
</tr>
<tr>
<td>
<code>dataDisks</code></br>
<em>
<a href="#onecloud.yunion.io/v1.VMDiskSpec">
[]VMDiskSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>DataDisks describes the configuration of data disks</p>
</td>
</tr>
<tr>
<td>
<code>networks</code></br>
<em>
<a href="#onecloud.yunion.io/v1.VMNetworkSpec">
[]VMNetworkSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
</tbody>
</table>
<h3 id="onecloud.yunion.io/v1.VirtualMachineSpec">VirtualMachineSpec
</h3>
<p>
(<em>Appears on:</em>
<a href="#onecloud.yunion.io/v1.VirtualMachine">VirtualMachine</a>)
</p>
<p>
<p>VirtualMachineSpec defines the desired state of VirtualMachine</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>DisplayIdenSpec</code></br>
<em>
<a href="#onecloud.yunion.io/v1.DisplayIdenSpec">
DisplayIdenSpec
</a>
</em>
</td>
<td>
<p>
(Members of <code>DisplayIdenSpec</code> are embedded into this type.)
</p>
<em>(Required)</em>
</td>
</tr>
<tr>
<td>
<code>vmConfig</code></br>
<em>
<a href="#onecloud.yunion.io/v1.VirtualMachineConfig">
VirtualMachineConfig
</a>
</em>
</td>
<td>
<em>(Required)</em>
</td>
</tr>
<tr>
<td>
<code>projectConfig</code></br>
<em>
<a href="#onecloud.yunion.io/v1.VMProjectSpec">
VMProjectSpec
</a>
</em>
</td>
<td>
<em>(Required)</em>
</td>
</tr>
<tr>
<td>
<code>VMPasswordSpec</code></br>
<em>
<a href="#onecloud.yunion.io/v1.VMPasswordSpec">
VMPasswordSpec
</a>
</em>
</td>
<td>
<p>
(Members of <code>VMPasswordSpec</code> are embedded into this type.)
</p>
<em>(Required)</em>
</td>
</tr>
<tr>
<td>
<code>secgroups</code></br>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
<tr>
<td>
<code>newEip</code></br>
<em>
<a href="#onecloud.yunion.io/v1.VMNewEipSpec">
VMNewEipSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>NewEip indicates that create a new EIP and bind it with VM.
It conflicts with Eip.</p>
</td>
</tr>
<tr>
<td>
<code>eip</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Eip indicates that bind the existing EIP with VM.
It conflicts with NewEip.</p>
</td>
</tr>
<tr>
<td>
<code>billDuration</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>BillDuration describes the duration of the annual and monthly billing type.
That length of BillDuration represents the billing type is &lsquo;prepaid&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>autoRenew</code></br>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>AutoRenew indicates whether to automatically renewal.</p>
</td>
</tr>
<tr>
<td>
<code>recreatePolicy</code></br>
<em>
<a href="#onecloud.yunion.io/v1.RecreatePolicy">
RecreatePolicy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
</tbody>
</table>
<h3 id="onecloud.yunion.io/v1.VirtualMachineStatus">VirtualMachineStatus
</h3>
<p>
(<em>Appears on:</em>
<a href="#onecloud.yunion.io/v1.VirtualMachine">VirtualMachine</a>)
</p>
<p>
<p>VirtualMachineStatus defines the observed state of VirtualMachine</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>phase</code></br>
<em>
<a href="#onecloud.yunion.io/v1.ResourcePhase">
ResourcePhase
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
Important: Run &ldquo;make&rdquo; to regenerate code after modifying this file</p>
</td>
</tr>
<tr>
<td>
<code>reason</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>A human readable message indicating details about why vm is in this phase.</p>
</td>
</tr>
<tr>
<td>
<code>externalInfo</code></br>
<em>
<a href="#onecloud.yunion.io/v1.VMInfo">
VMInfo
</a>
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
<tr>
<td>
<code>createTimes</code></br>
<em>
int32
</em>
</td>
<td>
<em>(Required)</em>
<p>CreateTimes record the continuous creation times.</p>
</td>
</tr>
</tbody>
</table>
<hr/>
<p><em>
Generated with <code>gen-crd-api-reference-docs</code>
on git commit <code>655cf96</code>.
</em></p>
