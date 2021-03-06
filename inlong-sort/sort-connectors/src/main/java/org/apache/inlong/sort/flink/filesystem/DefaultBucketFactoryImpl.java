/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements. See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package org.apache.inlong.sort.flink.filesystem;

import org.apache.flink.core.fs.Path;
import org.apache.flink.core.fs.RecoverableWriter;

import java.io.IOException;

/**
 * Copied from Apache Flink project, org.apache.flink.streaming.api.functions.sink.filesystem.DefaultBucketFactoryImpl.
 */
public class DefaultBucketFactoryImpl<IN, BucketID> implements
        BucketFactory<IN, BucketID> {

    private static final long serialVersionUID = 1L;

    @Override
    public Bucket<IN, BucketID> getNewBucket(
            final long dataFlowId,
            final RecoverableWriter fsWriter,
            final int subtaskIndex,
            final BucketID bucketId,
            final Path bucketPath,
            final long initialPartCounter,
            final PartFileWriter.PartFileFactory<IN, BucketID> partFileWriterFactory,
            final RollingPolicy<IN, BucketID> rollingPolicy) {

        return Bucket.getNew(
                dataFlowId,
                fsWriter,
                subtaskIndex,
                bucketId,
                bucketPath,
                initialPartCounter,
                partFileWriterFactory,
                rollingPolicy);
    }

    @Override
    public Bucket<IN, BucketID> restoreBucket(
            final long dataFlowId,
            final RecoverableWriter fsWriter,
            final int subtaskIndex,
            final long initialPartCounter,
            final PartFileWriter.PartFileFactory<IN, BucketID> partFileWriterFactory,
            final RollingPolicy<IN, BucketID> rollingPolicy,
            final BucketState<BucketID> bucketState) throws IOException {

        return Bucket.restore(
                dataFlowId,
                fsWriter,
                subtaskIndex,
                initialPartCounter,
                partFileWriterFactory,
                rollingPolicy,
                bucketState);
    }
}
