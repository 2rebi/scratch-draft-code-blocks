import {useCallback, useEffect, useRef} from "react";

export default function useInterval(callback: () => void, ms: number) {
    const intervalRef = useRef<number>();

    const start = useCallback(() => {
        if (intervalRef.current !== undefined || window === undefined) {
            return;
        }

        intervalRef.current = window.setInterval(callback, ms);
    }, [callback, ms]);

    const clear = useCallback(() => {
        const interval = intervalRef.current;
        if (interval === undefined || window === undefined) {
            return;
        }

        intervalRef.current = undefined;
        window.clearInterval(interval);
    }, []);

    useEffect(() => {
        start();
        return () => {
            clear()
        };
    }, [start, clear]);

    return { start, clear };
}
