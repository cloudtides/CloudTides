import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { map } from 'rxjs/operators';
import { environment } from '@tide-environments/environment';
import { VCD_URL_PATH } from '@tide-config/path';
import { LoginService } from '../login/login.service';
import toFixed from 'accounting-js/lib/toFixed.js';

@Injectable()
export class ResourceService {

  constructor(
    private readonly http: HttpClient,
    private readonly loginService: LoginService,
  ) {
  }

  private prefix = `${environment.apiPrefix}/computeResource`;

  async getList() {
    const list = await this.http.get<Item[]>(environment.apiPrefix + VCD_URL_PATH, {
      headers: {
        Authorization: `Bearer ${this.loginService.token}`,
      },
    }).toPromise();
    const usage: Item[] = [];
    for (const resource of list) {
      const rawUsage = await this.http.get<ItemUsage>(`${environment.apiPrefix}/usage/${resource.id}`, {
        headers: {
          Authorization: `Bearer ${this.loginService.token}`,
        },
      }).toPromise();
      const resourceItem: Item = {
        id: resource.id,
        vcdId: resource.vcdId,
        name: rawUsage.name,
        organization: resource.organization,
        vendor: resource.vendor,
        cpu: rawUsage.totalCPU / 1000,
        mem: rawUsage.totalRAM / 1024,
        disk: rawUsage.totalDisk / 1024,
        usage: {
          'cpu%': toFixed(rawUsage.percentCPU * 100, 2),
          'mem%': toFixed(rawUsage.percentRAM * 100, 2),
          'disk%': toFixed(rawUsage.percentDisk * 100, 2),
        },
      };
      usage.push(resourceItem);
    }
    return usage;
  }

  addItem(payload: ItemPayload) {
    const body = {
      ...payload,
      policy: 0,
    };
    return this.http.post<any>(environment.apiPrefix + VCD_URL_PATH, body, {
      headers: {
        Authorization: `Bearer ${this.loginService.token}`,
      },
    }).toPromise().then(() => {
      return Promise.resolve();
    }, (errResp) => {
      return Promise.reject(`HTTP ${errResp.status}: ${errResp.error.message}`);
    });
  }

  editItem(id: string, payload: ItemPayload) {
    return this.http.put<ItemDTO>(`${this.prefix}/${id}`, payload).pipe(
      map(mapItem),
    );
  }

  async removeItem(id: string) {
    await this.http.delete<any>(environment.apiPrefix + `/resource/vcd/` + id, {
      headers: {
        Authorization: `Bearer ${this.loginService.token}`,
    }, }).toPromise().then(
      () => {
        return Promise.resolve();
      }, (errResp) => {
        return Promise.reject(`${errResp.message}`);
      },
    );
  }

  async contributeResource(id: string): Promise<ContributeResp> {
    let response = null;
    await this.http.put(environment.apiPrefix + `/resource/contribute/${id}`, null, {
      headers: {
      Authorization: `Bearer ${this.loginService.token}`,
    }, }).toPromise().then(
      (resp) => {
        response = resp;
        return Promise.resolve();
      }, (errResp) => {
        return Promise.reject(`${errResp.message}`);
      },
    );
    return response;
  }

  async activateResource(id: string): Promise<ActivateResp> {
    let response = null;
    await this.http.put<ActivateResp>(environment.apiPrefix + `/resource/activate/${id}`, null, {
      headers: {
        Authorization: `Bearer ${this.loginService.token}`,
      },
    }).toPromise().then(
      (resp) => {
        response = resp;
        return Promise.resolve();
      }, (errResp) => {
        return Promise.reject(`${errResp.message}`);
      },
    );
    return response;
  }
}

// Raw
interface ItemUsage {
  currentCPU: number;
  totalCPU: number;
  currentRAM: number;
  totalRAM: number;
  currentDisk: number;
  totalDisk: number;
  percentCPU: number;
  percentRAM: number;
  percentDisk: number;
  name: string;
}

interface ItemDTO {
  id: string;
  vcdId: string;
  name: string;
  organization: string;
  vendor: string;
  // unit: GHz
  cpu: number;
  // unit: GB
  mem: number;
  // unit: GB
  disk: number;
  usage: {
    'cpu%': number;
    'mem%': number;
    'disk%': number;
  }
}

interface ContributeResp {
  message: string;
  contributed: boolean;
}

interface ActivateResp {
  message: string;
  activated: boolean;
}

function mapList(raw: ItemDTO[]): Item[] {
  return raw.map(mapItem);
}

function mapItem(raw: ItemDTO): Item {
  return raw;
}

// UI
export interface ItemPayload {
  datacenter: string;
  name: string;
  org: string;
  network: string;
  catalog: string;
  username: string,
  password: string,
}

export type Item = ItemDTO;
