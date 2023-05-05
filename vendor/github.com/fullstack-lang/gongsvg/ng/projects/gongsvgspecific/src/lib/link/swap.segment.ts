import { Segment, createPoint } from './draw.segments';

export function swapSegment(segment: Segment): Segment {

    let res = structuredClone(segment)

    res.StartPoint = segment.EndPoint
    res.EndPoint = segment.StartPoint

    return res

}