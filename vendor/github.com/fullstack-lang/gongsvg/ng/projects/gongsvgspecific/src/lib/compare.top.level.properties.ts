


export function compareTopLevelProperties<T extends object>(instance1: T, instance2: T): boolean {
    const keys1 = Object.keys(instance1) as (keyof T)[];
    const keys2 = Object.keys(instance2) as (keyof T)[];

    if (keys1.length !== keys2.length) {
        return false;
    }

    for (let i = 0; i < keys1.length; i++) {
        const key = keys1[i];

        if (instance1[key] !== instance2[key]) {
            return false;
        }
    }

    return true;
}