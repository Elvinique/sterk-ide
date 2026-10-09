const URL = 'http://127.0.0.1:4187';
export type Entry = {name:string;path:string;dir:boolean};
export type Goal = {id:string;objective:string;criteria:string;status:string;created:string;evidence:string[]};
export function token(){return sessionStorage.getItem('sterk-token')||''}
export async function api<T>(path:string,options:RequestInit={}):Promise<T>{const response=await fetch(URL+path,{...options,headers:{'X-Sterk-Token':token(),'Content-Type':'application/json',...options.headers}});const result=await response.json();if(!response.ok)throw new Error(result.error||`HTTP ${response.status}`);return result as T}
