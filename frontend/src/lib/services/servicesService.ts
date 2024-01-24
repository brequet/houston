export interface Service {
  name: string;
  description: string;
  status: number;
  processId: number;
}

export const servicesService = {
  async getServices(): Promise<Service[]> {
    const response = await fetch(`${import.meta.env.VITE_SERVER_URL}/services`);
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    const scripts = (await response.json()) as Service[];
    return scripts;
  },
};

export type ServicesService = typeof servicesService;
