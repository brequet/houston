export interface FileDetails {
    name: string;
    type: string;
    modTime: Date;
    extension: string;
}

export interface ScriptExecutionResponse {
    output: string;
    error?: string;
    scriptName: string;
}

export const scriptsService = {
    async getScripts(): Promise<FileDetails[]> {
        const response = await fetch(`${import.meta.env.VITE_SERVER_URL}/api/scripts`);
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        const scripts = await response.json() as FileDetails[];
        return scripts;
    },
    async executeScriptByName(scriptName: string): Promise<ScriptExecutionResponse> {
        const response = await fetch(`${import.meta.env.VITE_SERVER_URL}/api/script/execute`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ scriptName })
        });

        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }

        const data: ScriptExecutionResponse = await response.json();
        return data;
    }
}

export type ScriptsService = typeof scriptsService;
